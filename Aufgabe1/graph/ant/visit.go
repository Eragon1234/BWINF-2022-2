package ant

import (
	"BWINF/Aufgabe1/graph"
	"BWINF/Aufgabe1/vector"
	"BWINF/pkg/set"
	"BWINF/pkg/slice"
	"sort"
	"sync"
)

func VisitAll(cfg Config, g graph.WeightedGraph[vector.Coordinate, graph.DistanceAngle]) []graph.Edge[PheromoneDistanceAngle, vector.Coordinate] {
	pheromoneGraph := TransformGraph(g)

	// we add a start vertex to the graph that has length 0 to all other vertices
	// this is necessary because the ant colony optimization algorithm needs a start vertex
	pheromoneGraph.AddVertex("start", vector.Coordinate{})
	for _, v := range pheromoneGraph.Vertices {
		pheromoneGraph.AddEdge(pheromoneGraph.Vertices["start"], v, PheromoneDistanceAngle{
			Pheromone:     cfg.PheromoneEvaporation,
			DistanceAngle: graph.DistanceAngle{Distance: 0, Angle: 0},
		})
	}
	var shortestPath []graph.Edge[PheromoneDistanceAngle, vector.Coordinate]
	updatePheromone(cfg, pheromoneGraph, [][]graph.Edge[PheromoneDistanceAngle, vector.Coordinate]{transformEdges(graph.VisitAllShortestEdge(g))})

	ants := slice.MakeFunc(cfg.NumOfAnts, func(i int) ant {
		return *newAnt(cfg.PheromoneWeight, cfg.DistanceWeight)
	})

	var wg sync.WaitGroup
	for i := 0; i < cfg.NumOfIterations; i++ {
		fmt.Fprintf(os.Stderr, "iteration, bestLength: %d %f\n", i, LengthOfPheromonePath(shortestPath))
		resultChan := make(chan []graph.Edge[PheromoneDistanceAngle, vector.Coordinate], cfg.NumOfAnts)
		wg.Add(cfg.NumOfAnts)
		for _, ant := range ants {
			go ant.Run(&wg, pheromoneGraph, resultChan)
		}
		wg.Wait()
		close(resultChan)

		var newPaths [][]graph.Edge[PheromoneDistanceAngle, vector.Coordinate]
		for path := range resultChan {
			newPaths = append(newPaths, path)
		}

		sort.Slice(newPaths, func(i, j int) bool {
			return LengthOfPheromonePath(newPaths[i]) < LengthOfPheromonePath(newPaths[j])
		})

		if len(newPaths) == 0 {
			continue
		}

		if shortestPath == nil || LengthOfPheromonePath(newPaths[0]) < LengthOfPheromonePath(shortestPath) {
			shortestPath = newPaths[0]
		}

		updatePheromone(cfg, pheromoneGraph, newPaths[:cfg.Elite])
		print("\033[F\033[2K")
	}
	return shortestPath
}

func updatePheromone(cfg Config, pheromoneGraph *graph.WeightedGraph[vector.Coordinate, PheromoneDistanceAngle], paths [][]graph.Edge[PheromoneDistanceAngle, vector.Coordinate]) {
	var pathSets []set.Set[graph.Edge[PheromoneDistanceAngle, vector.Coordinate]]
	for _, path := range paths {
		pathSet := set.FromSlice(path)
		pathSets = append(pathSets, *pathSet)
	}
	for _, v := range pheromoneGraph.Vertices {
		for _, e := range pheromoneGraph.GetEdges(v) {
			var amount float64

			for i, pathSet := range pathSets {
				if pathSet.Contains(e) {
					amount += cfg.PheromoneAmount / LengthOfPheromonePath(paths[i])
				}
			}

			pheromone := e.Weight.Pheromone

			pheromoneGraph.UpdateEdge(e.From, e.To, PheromoneDistanceAngle{
				Pheromone:     (1-cfg.PheromoneEvaporation)*pheromone + amount,
				DistanceAngle: e.Weight.DistanceAngle,
			})
		}
	}
}
