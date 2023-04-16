package ant

import (
	"BWINF/Aufgabe1/graph"
	"BWINF/Aufgabe1/vector"
	"BWINF/pkg/slice"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"sync"
)

func VisitAllAntColonyOptimization(cfg Config, g graph.WeightedGraph[vector.Coordinate, graph.DistanceAngle]) []graph.Edge[PheromoneDistanceAngle, vector.Coordinate] {
	pheromoneGraph := transformGraph(g)

	var totalEdgeLength float64
	for _, v := range g.Vertices {
		for _, e := range g.GetEdges(v) {
			totalEdgeLength += e.Weight.Distance
		}
	}

	// we add a start vertex to the graph that has length 0 to all other vertices
	// this is necessary because the ant colony optimization algorithm needs a start vertex
	pheromoneGraph.AddVertex("start", vector.Coordinate{})
	for _, v := range pheromoneGraph.Vertices {
		pheromoneGraph.AddEdge(pheromoneGraph.Vertices["start"], v, PheromoneDistanceAngle{
			Pheromone:     0,
			DistanceAngle: graph.DistanceAngle{Distance: 0, Angle: 0},
		})
	}
	shortestPath := transformEdges(graph.VisitAllShortestEdge(g))
	updatePheromone(cfg, totalEdgeLength, pheromoneGraph, shortestPath)

	shortestPath = pheromoneGraph.GetEdges(pheromoneGraph.Vertices["210.000000 30.000000"])

	ants := slice.MakeFunc(cfg.NumOfAnts, func(i int) Ant {
		return *NewAnt(i, cfg)
	})

	var wg sync.WaitGroup
	for i := 0; i < cfg.NumOfIterations; i++ {
		fmt.Fprintln(os.Stderr, "iteration, bestLength:", i, LengthOfPheromonePath(shortestPath))

		result := make(chan antResult, cfg.NumOfAnts)
		wg.Add(len(ants))
		for _, ant := range ants {
			go ant.Run(&wg, pheromoneGraph, result, LengthOfPheromonePath(shortestPath))
		}
		wg.Wait()
		close(result)

		var results []antResult
		for path := range result {
			results = append(results, path)
		}

		sort.Slice(results, func(i, j int) bool {
			return LengthOfPheromonePath(results[i].Path) < LengthOfPheromonePath(results[j].Path)
		})

		var eliteAnts []Ant

		eliteIndex := int(float64(len(results)) * cfg.EliteProportion)
		for _, antResult := range results[:eliteIndex] {
			eliteAnts = append(eliteAnts, ants[antResult.Index])
			updatePheromone(cfg, totalEdgeLength, pheromoneGraph, antResult.Path)
			if LengthOfPheromonePath(antResult.Path) < LengthOfPheromonePath(shortestPath) {
				shortestPath = antResult.Path
			}
		}
		for _, loserAnt := range ants[eliteIndex:] {
			randomEliteAnt := eliteAnts[rand.Intn(len(eliteAnts))]
			loserAnt.Optimize(randomEliteAnt)
		}

		fmt.Fprint(os.Stderr, "\u001B[1A\u001B[K")
	}
	return shortestPath
}
