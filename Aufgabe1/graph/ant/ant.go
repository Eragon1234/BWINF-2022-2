package ant

import (
	"BWINF/Aufgabe1/graph"
	"BWINF/Aufgabe1/vector"
	"BWINF/pkg"
	"BWINF/pkg/set"
	"BWINF/pkg/slice"
	"fmt"
	"math"
	"sync"
)

type ant struct {
	PheromoneWeight float64
	DistanceWeight  float64
}

func newAnt(PheromoneWeight float64, DistanceWeight float64) *ant {
	return &ant{
		PheromoneWeight: PheromoneWeight,
		DistanceWeight:  DistanceWeight,
	}
}

func (a *ant) Run(wg *sync.WaitGroup, g *graph.WeightedGraph[vector.Coordinate, PheromoneDistanceAngle], result chan<- []graph.Edge[PheromoneDistanceAngle, vector.Coordinate]) {
	defer wg.Done()

	var visited set.Set[*graph.Vertex[vector.Coordinate]]
	visited.Add(g.Vertices["start"])

	path := make([]graph.Edge[PheromoneDistanceAngle, vector.Coordinate], 0, len(g.Vertices))
	for visited.Size() < len(g.Vertices) {
		var edges []graph.Edge[PheromoneDistanceAngle, vector.Coordinate]
		if len(path) == 0 {
			edges = g.GetEdges(g.Vertices["start"])
			edges = slice.FilterFunc(edges, func(e graph.Edge[PheromoneDistanceAngle, vector.Coordinate]) bool {
				return visited.Contains(e.To)
			})
		} else {
			current := path[len(path)-1]
			if current.To == nil {
				fmt.Printf("current.To is nil path: %v\n", path)
			}
			edges = g.GetEdges(current.To)
			edges = slice.FilterFunc(edges, func(e graph.Edge[PheromoneDistanceAngle, vector.Coordinate]) bool {
				return visited.Contains(e.To)
			})
			turnAngleFiltered := slice.FilterFunc(edges, func(e graph.Edge[PheromoneDistanceAngle, vector.Coordinate]) bool {
				return vector.TurnAngle(current.Weight.DistanceAngle.Angle, e.Weight.DistanceAngle.Angle) > 90
			})
			if len(turnAngleFiltered) > 0 {
				edges = turnAngleFiltered
			}
		}
		edge := a.chooseNextEdge(edges)
		visited.Add(edge.To)
		path = append(path, edge)
	}
	result <- path
}

func (a *ant) chooseNextEdge(edges []graph.Edge[PheromoneDistanceAngle, vector.Coordinate]) graph.Edge[PheromoneDistanceAngle, vector.Coordinate] {
	weights := make([]float64, len(edges))
	for i, e := range edges {
		weights[i] = math.Pow(1/e.Weight.DistanceAngle.Distance, a.DistanceWeight) * math.Pow(math.Max(e.Weight.Pheromone, 1), a.PheromoneWeight)
	}
	edge := edges[pkg.ChooseWeighted(weights)]
	return edge
}

func LengthOfPheromonePath(pheromonePath []graph.Edge[PheromoneDistanceAngle, vector.Coordinate]) float64 {
	totalDistance := 0.0
	for _, e := range pheromonePath {
		totalDistance += e.Weight.DistanceAngle.Distance
	}
	return totalDistance
}

type PheromoneDistanceAngle struct {
	Pheromone     float64
	DistanceAngle graph.DistanceAngle
}
