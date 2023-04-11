package graph

import (
	"BWINF/Aufgabe1/vector"
	"BWINF/utils"
	"BWINF/utils/slice"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"sync"
)

var numOfAnts = runtime.NumCPU() * 32
var numOfIterations = 1_000

var randomChance = utils.EaseOut{
	From: 1,
	To:   0.3,
}
var pheromoneWeight = utils.EaseOut{
	From: 1,
	To:   5,
}
var distanceWeight = utils.EaseOut{
	From: 3,
	To:   1,
}
var pheromoneDecreasement = utils.EaseOut{
	From: 0.7,
	To:   0.9,
}

var eliteProportion = 0.1
var cutoffWhenLongerThanShortest = utils.EaseOut{
	From: 100,
	To:   1.5,
}

var t float64

func VisitAllAntColonyOptimization(g WeightedGraph[vector.Coordinate, DistanceAngle]) []Edge[pheromoneDistanceAngle, vector.Coordinate] {
	pheromoneGraph := transformGraph(g)

	// we add a start vertex to the graph that has length 0 to all other vertices
	// this is necessary because the ant colony optimization algorithm needs a start vertex
	pheromoneGraph.AddVertex("start", vector.Coordinate{})
	for _, v := range pheromoneGraph.Vertices {
		pheromoneGraph.AddEdge(pheromoneGraph.Vertices["start"], v, pheromoneDistanceAngle{
			Pheromone:     0,
			DistanceAngle: DistanceAngle{Distance: 0, Angle: 0},
		})
	}
	shortestPath := transformEdges(VisitAllShortestEdge(g))
	updatePheromone(pheromoneGraph, shortestPath)

	//shortestPath = pheromoneGraph.GetEdges(pheromoneGraph.Vertices["210.000000 30.000000"])
	var wg sync.WaitGroup

	for i := 0; i < numOfIterations; i++ {
		t = float64(i) / float64(numOfIterations)
		fmt.Fprintln(os.Stderr, "iteration, bestLength:", i, LengthOfPheromonePath(shortestPath))

		result := make(chan []Edge[pheromoneDistanceAngle, vector.Coordinate], numOfAnts)
		wg.Add(numOfAnts)
		for i := 0; i < numOfAnts; i++ {
			go AntDoPath(&wg, *pheromoneGraph, result, LengthOfPheromonePath(shortestPath))
		}

		wg.Wait()
		close(result)

		var newPaths [][]Edge[pheromoneDistanceAngle, vector.Coordinate]
		for path := range result {
			newPaths = append(newPaths, path)
		}
		if newPaths != nil {
			sort.Slice(newPaths, func(i, j int) bool {
				return LengthOfPheromonePath(newPaths[i]) < LengthOfPheromonePath(newPaths[j])
			})
			for _, path := range newPaths[:int(float64(len(newPaths))*eliteProportion)] {
				updatePheromone(pheromoneGraph, path)
				if LengthOfPheromonePath(path) < LengthOfPheromonePath(shortestPath) {
					shortestPath = path
				}
			}
		}

		fmt.Fprint(os.Stderr, "\u001B[1A\u001B[K")
	}
	return shortestPath
}

func AntDoPath(wg *sync.WaitGroup, g WeightedGraph[vector.Coordinate, pheromoneDistanceAngle], result chan<- []Edge[pheromoneDistanceAngle, vector.Coordinate], shortestLength float64) {
	defer wg.Done()

	path := make([]Edge[pheromoneDistanceAngle, vector.Coordinate], 0, len(g.Vertices))
	var visited utils.Set[*Vertex[vector.Coordinate]]
	var current = Edge[pheromoneDistanceAngle, vector.Coordinate]{
		From: g.Vertices["start"],
		To:   g.Vertices["start"],
		Weight: pheromoneDistanceAngle{
			Pheromone:     1,
			DistanceAngle: DistanceAngle{Distance: 0, Angle: 0},
		},
	}
	visited.Add(current.To)
	for visited.Size() < len(g.Vertices) {
		if LengthOfPheromonePath(path) > shortestLength*cutoffWhenLongerThanShortest.Get(t) {
			break
		}
		edges := g.GetEdges(current.To)
		edges = slice.FilterFunc(edges, func(e Edge[pheromoneDistanceAngle, vector.Coordinate]) bool {
			return visited.Contains(e.To)
		})
		turnAngleFiltered := slice.FilterFunc(edges, func(e Edge[pheromoneDistanceAngle, vector.Coordinate]) bool {
			return vector.TurnAngle(current.Weight.DistanceAngle.Angle, e.Weight.DistanceAngle.Angle) > 90
		})
		if len(turnAngleFiltered) > 0 {
			edges = turnAngleFiltered
		}
		edge := chooseNextEdge(edges)
		visited.Add(edge.To)
		current = edge
		path = append(path, edge)
	}
	if visited.Size() == len(g.Vertices) {
		result <- path
	}
}

func chooseNextEdge(edges []Edge[pheromoneDistanceAngle, vector.Coordinate]) Edge[pheromoneDistanceAngle, vector.Coordinate] {
	q := rand.Float64()

	if q > randomChance.Get(t) {
		return chooseNextEdgeWeighted(edges)
	} else {
		n := rand.Intn(len(edges))
		return edges[n]
	}
}

func chooseNextEdgeWeighted(edges []Edge[pheromoneDistanceAngle, vector.Coordinate]) Edge[pheromoneDistanceAngle, vector.Coordinate] {
	weights := make([]float64, len(edges))
	for i, e := range edges {
		weights[i] = e.Weight.Pheromone * pheromoneWeight.Get(t) / (e.Weight.DistanceAngle.Distance * distanceWeight.Get(t))
	}
	n := utils.ChooseWeighted(weights)
	return edges[n]
}

func updatePheromone(g *WeightedGraph[vector.Coordinate, pheromoneDistanceAngle], shortestPath []Edge[pheromoneDistanceAngle, vector.Coordinate]) {
	totalDistance := LengthOfPheromonePath(shortestPath)
	for _, v := range g.Vertices {
		for _, e := range g.GetEdges(v) {
			g.UpdateEdge(e.From, e.To, pheromoneDistanceAngle{
				Pheromone:     e.Weight.Pheromone * pheromoneDecreasement.Get(t),
				DistanceAngle: e.Weight.DistanceAngle,
			})
		}
	}
	for _, e := range shortestPath {
		newPheromone := e.Weight.Pheromone * 1000 / totalDistance
		g.UpdateEdge(e.From, e.To, pheromoneDistanceAngle{
			Pheromone:     newPheromone,
			DistanceAngle: e.Weight.DistanceAngle,
		})
	}
}

func LengthOfPheromonePath(pheromonePath []Edge[pheromoneDistanceAngle, vector.Coordinate]) float64 {
	totalDistance := 0.0
	for _, e := range pheromonePath {
		totalDistance += e.Weight.DistanceAngle.Distance
	}
	return totalDistance
}

func transformGraph(g WeightedGraph[vector.Coordinate, DistanceAngle]) *WeightedGraph[vector.Coordinate, pheromoneDistanceAngle] {
	pheromoneGraph := NewWeightedGraph[vector.Coordinate, pheromoneDistanceAngle]()
	for _, v := range g.Vertices {
		pheromoneGraph.AddVertex(v.Name, v.Value)
	}
	for _, v := range g.Vertices {
		for _, e := range g.GetEdges(v) {
			pheromoneGraph.AddEdge(pheromoneGraph.Vertices[e.From.Name], pheromoneGraph.Vertices[e.To.Name], pheromoneDistanceAngle{
				Pheromone:     0,
				DistanceAngle: e.Weight,
			})
		}
	}
	return pheromoneGraph
}

func transformEdges(edges []Edge[DistanceAngle, vector.Coordinate]) []Edge[pheromoneDistanceAngle, vector.Coordinate] {
	pheromoneEdges := make([]Edge[pheromoneDistanceAngle, vector.Coordinate], len(edges))
	for i, e := range edges {
		pheromoneEdges[i] = Edge[pheromoneDistanceAngle, vector.Coordinate]{
			From:   e.From,
			To:     e.To,
			Weight: pheromoneDistanceAngle{Pheromone: 0, DistanceAngle: e.Weight},
		}
	}
	return pheromoneEdges
}

type pheromoneDistanceAngle struct {
	Pheromone     float64
	DistanceAngle DistanceAngle
}
