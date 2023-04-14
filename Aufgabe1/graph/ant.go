package graph

import (
	"BWINF/Aufgabe1/vector"
	"BWINF/utils"
	"BWINF/utils/slice"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"sync"
)

type Config struct {
	NumOfAnts                    int
	NumOfIterations              int
	RandomChance                 float64
	PheromoneWeight              float64
	DistanceWeight               float64
	PheromoneDecreasement        float64
	EliteProportion              float64
	CutoffWhenLongerThanShortest float64
}

func VisitAllAntColonyOptimization(cfg Config, g WeightedGraph[vector.Coordinate, DistanceAngle]) []Edge[PheromoneDistanceAngle, vector.Coordinate] {
	pheromoneGraph := transformGraph(g)

	// we add a start vertex to the graph that has length 0 to all other vertices
	// this is necessary because the ant colony optimization algorithm needs a start vertex
	pheromoneGraph.AddVertex("start", vector.Coordinate{})
	for _, v := range pheromoneGraph.Vertices {
		pheromoneGraph.AddEdge(pheromoneGraph.Vertices["start"], v, PheromoneDistanceAngle{
			Pheromone:     0,
			DistanceAngle: DistanceAngle{Distance: 0, Angle: 0},
		})
	}
	shortestPath := transformEdges(VisitAllShortestEdge(g))
	updatePheromone(cfg, pheromoneGraph, shortestPath)

	shortestPath = pheromoneGraph.GetEdges(pheromoneGraph.Vertices["210.000000 30.000000"])
	var wg sync.WaitGroup

	for i := 0; i < cfg.NumOfIterations; i++ {
		fmt.Fprintln(os.Stderr, "iteration, bestLength:", i, LengthOfPheromonePath(shortestPath))

		result := make(chan []Edge[PheromoneDistanceAngle, vector.Coordinate], cfg.NumOfAnts)
		wg.Add(cfg.NumOfAnts)
		for i := 0; i < cfg.NumOfAnts; i++ {
			go AntDoPath(cfg, &wg, result, LengthOfPheromonePath(shortestPath), *pheromoneGraph)
		}

		wg.Wait()
		close(result)

		var newPaths [][]Edge[PheromoneDistanceAngle, vector.Coordinate]
		for path := range result {
			newPaths = append(newPaths, path)
		}
		if newPaths != nil {
			sort.Slice(newPaths, func(i, j int) bool {
				return LengthOfPheromonePath(newPaths[i]) < LengthOfPheromonePath(newPaths[j])
			})
			for _, path := range newPaths[:int(float64(len(newPaths))*cfg.EliteProportion)] {
				updatePheromone(cfg, pheromoneGraph, path)
				if LengthOfPheromonePath(path) < LengthOfPheromonePath(shortestPath) {
					shortestPath = path
				}
			}
		}

		fmt.Fprint(os.Stderr, "\u001B[1A\u001B[K")
	}
	return shortestPath
}

func AntDoPath(cfg Config, wg *sync.WaitGroup, result chan<- []Edge[PheromoneDistanceAngle, vector.Coordinate], shortestLength float64, g WeightedGraph[vector.Coordinate, PheromoneDistanceAngle]) {
	defer wg.Done()

	path := make([]Edge[PheromoneDistanceAngle, vector.Coordinate], 0, len(g.Vertices))
	var visited utils.Set[*Vertex[vector.Coordinate]]
	var current = Edge[PheromoneDistanceAngle, vector.Coordinate]{
		From: g.Vertices["start"],
		To:   g.Vertices["start"],
		Weight: PheromoneDistanceAngle{
			Pheromone:     1,
			DistanceAngle: DistanceAngle{Distance: 0, Angle: 0},
		},
	}
	visited.Add(current.To)
	for visited.Size() < len(g.Vertices) {
		if LengthOfPheromonePath(path) > shortestLength*cfg.CutoffWhenLongerThanShortest {
			break
		}
		edges := g.GetEdges(current.To)
		edges = slice.FilterFunc(edges, func(e Edge[PheromoneDistanceAngle, vector.Coordinate]) bool {
			return visited.Contains(e.To)
		})
		turnAngleFiltered := slice.FilterFunc(edges, func(e Edge[PheromoneDistanceAngle, vector.Coordinate]) bool {
			return vector.TurnAngle(current.Weight.DistanceAngle.Angle, e.Weight.DistanceAngle.Angle) > 90
		})
		if len(turnAngleFiltered) > 0 {
			edges = turnAngleFiltered
		}
		edge := chooseNextEdge(cfg, edges)
		visited.Add(edge.To)
		current = edge
		path = append(path, edge)
	}
	if visited.Size() == len(g.Vertices) {
		result <- path
	}
}

func chooseNextEdge(cfg Config, edges []Edge[PheromoneDistanceAngle, vector.Coordinate]) Edge[PheromoneDistanceAngle, vector.Coordinate] {
	q := rand.Float64()

	if q > cfg.RandomChance {
		return chooseNextEdgeWeighted(cfg, edges)
	} else {
		n := rand.Intn(len(edges))
		return edges[n]
	}
}

func chooseNextEdgeWeighted(cfg Config, edges []Edge[PheromoneDistanceAngle, vector.Coordinate]) Edge[PheromoneDistanceAngle, vector.Coordinate] {
	weights := make([]float64, len(edges))
	for i, e := range edges {
		weights[i] = e.Weight.Pheromone * cfg.PheromoneWeight / (e.Weight.DistanceAngle.Distance * cfg.DistanceWeight)
	}
	n := utils.ChooseWeighted(weights)
	return edges[n]
}

func updatePheromone(cfg Config, g *WeightedGraph[vector.Coordinate, PheromoneDistanceAngle], shortestPath []Edge[PheromoneDistanceAngle, vector.Coordinate]) {
	totalDistance := LengthOfPheromonePath(shortestPath)
	for _, v := range g.Vertices {
		for _, e := range g.GetEdges(v) {
			g.UpdateEdge(e.From, e.To, PheromoneDistanceAngle{
				Pheromone:     e.Weight.Pheromone * cfg.PheromoneDecreasement,
				DistanceAngle: e.Weight.DistanceAngle,
			})
		}
	}
	for _, e := range shortestPath {
		newPheromone := e.Weight.Pheromone * 1000 / totalDistance
		g.UpdateEdge(e.From, e.To, PheromoneDistanceAngle{
			Pheromone:     newPheromone,
			DistanceAngle: e.Weight.DistanceAngle,
		})
	}
}

func LengthOfPheromonePath(pheromonePath []Edge[PheromoneDistanceAngle, vector.Coordinate]) float64 {
	totalDistance := 0.0
	for _, e := range pheromonePath {
		totalDistance += e.Weight.DistanceAngle.Distance
	}
	return totalDistance
}

func transformGraph(g WeightedGraph[vector.Coordinate, DistanceAngle]) *WeightedGraph[vector.Coordinate, PheromoneDistanceAngle] {
	pheromoneGraph := NewWeightedGraph[vector.Coordinate, PheromoneDistanceAngle]()
	for _, v := range g.Vertices {
		pheromoneGraph.AddVertex(v.Name, v.Value)
	}
	for _, v := range g.Vertices {
		for _, e := range g.GetEdges(v) {
			pheromoneGraph.AddEdge(pheromoneGraph.Vertices[e.From.Name], pheromoneGraph.Vertices[e.To.Name], PheromoneDistanceAngle{
				Pheromone:     0,
				DistanceAngle: e.Weight,
			})
		}
	}
	return pheromoneGraph
}

func transformEdges(edges []Edge[DistanceAngle, vector.Coordinate]) []Edge[PheromoneDistanceAngle, vector.Coordinate] {
	pheromoneEdges := make([]Edge[PheromoneDistanceAngle, vector.Coordinate], len(edges))
	for i, e := range edges {
		pheromoneEdges[i] = Edge[PheromoneDistanceAngle, vector.Coordinate]{
			From:   e.From,
			To:     e.To,
			Weight: PheromoneDistanceAngle{Pheromone: 0, DistanceAngle: e.Weight},
		}
	}
	return pheromoneEdges
}

type PheromoneDistanceAngle struct {
	Pheromone     float64
	DistanceAngle DistanceAngle
}
