package ant

import (
	"BWINF/Aufgabe1/graph"
	"BWINF/Aufgabe1/vector"
	"BWINF/pkg"
	"BWINF/pkg/set"
	"BWINF/pkg/slice"
	"fmt"
	"math/rand"
	"sync"
)

type antResult struct {
	Path  []graph.Edge[PheromoneDistanceAngle, vector.Coordinate]
	Index int
}

type Ant struct {
	Index int
	cfg   Config
}

func NewAnt(i int, cfg Config) *Ant {
	return &Ant{Index: i, cfg: cfg}
}

func (a *Ant) Run(wg *sync.WaitGroup, g *graph.WeightedGraph[vector.Coordinate, PheromoneDistanceAngle], result chan<- antResult, shortestLength float64) {
	defer wg.Done()

	var visited set.Set[*graph.Vertex[vector.Coordinate]]
	visited.Add(g.Vertices["start"])

	path := make([]graph.Edge[PheromoneDistanceAngle, vector.Coordinate], 0, len(g.Vertices))
	for visited.Size() < len(g.Vertices) {
		if LengthOfPheromonePath(path) > shortestLength*a.cfg.CutoffWhenLongerThanShortest {
			break
		}
		var edges []graph.Edge[PheromoneDistanceAngle, vector.Coordinate]
		if len(path) == 0 {
			edges = g.GetEdges(g.Vertices["start"])
			edges = slice.FilterFunc(edges, func(e graph.Edge[PheromoneDistanceAngle, vector.Coordinate]) bool {
				return visited.Contains(e.To)
			})
		} else {
			current := path[len(path)-1]
			if current.To == nil {
				fmt.Printf("path: %v\n", path)
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
	if visited.Size() == len(g.Vertices) {
		result <- antResult{
			Path:  path,
			Index: a.Index,
		}
	}
}

func (a *Ant) Optimize(otherAnt Ant) {
	switch rand.Intn(8) {
	case 0:
		a.cfg = otherAnt.cfg
	case 1, 2:
		a.RandomizeRandomParameter()
	case 3, 4:
		a.TakeRandomParameter(otherAnt)
	case 5, 6:
		a.ModifyRandomParameter()
	case 7:
		a.RandomizeAllParameters()
	}
}

func (a *Ant) ModifyRandomParameter() {
	switch rand.Intn(4) {
	case 0:
		a.cfg.RandomChance += rand.Float64() - 0.5
		a.cfg.RandomChance = pkg.Clamp(0, a.cfg.RandomChance, 1)
	case 1:
		a.cfg.PheromoneWeight += rand.Float64() - 0.5
		a.cfg.PheromoneWeight = pkg.Clamp(0, a.cfg.PheromoneWeight, 1)
	case 2:
		a.cfg.DistanceWeight += rand.Float64() - 0.5
		a.cfg.DistanceWeight = pkg.Clamp(0, a.cfg.DistanceWeight, 1)
	case 3:
		a.cfg.CutoffWhenLongerThanShortest += (rand.Float64() - 0.5) * 10
		a.cfg.CutoffWhenLongerThanShortest = pkg.Clamp(2, a.cfg.CutoffWhenLongerThanShortest, 10)
	}
}
func (a *Ant) TakeRandomParameter(otherAnt Ant) {
	switch rand.Intn(4) {
	case 0:
		a.cfg.RandomChance = otherAnt.cfg.RandomChance
	case 1:
		a.cfg.PheromoneWeight = otherAnt.cfg.PheromoneWeight
	case 2:
		a.cfg.DistanceWeight = otherAnt.cfg.DistanceWeight
	case 3:
		a.cfg.CutoffWhenLongerThanShortest = otherAnt.cfg.CutoffWhenLongerThanShortest
	}
}

func (a *Ant) RandomizeRandomParameter() {
	switch rand.Intn(4) {
	case 0:
		a.cfg.RandomChance = rand.Float64()
	case 1:
		a.cfg.PheromoneWeight = rand.Float64()
	case 2:
		a.cfg.DistanceWeight = rand.Float64()
	case 3:
		a.cfg.CutoffWhenLongerThanShortest = rand.Float64() * 10
		a.cfg.CutoffWhenLongerThanShortest = pkg.Clamp(2, a.cfg.CutoffWhenLongerThanShortest, 10)
	}
}

func (a *Ant) RandomizeAllParameters() {
	a.cfg.RandomChance = rand.Float64()
	a.cfg.PheromoneWeight = rand.Float64()
	a.cfg.DistanceWeight = rand.Float64()
	a.cfg.CutoffWhenLongerThanShortest = rand.Float64() * 10
	a.cfg.CutoffWhenLongerThanShortest = pkg.Clamp(2, a.cfg.CutoffWhenLongerThanShortest, 10)
}

func (a *Ant) chooseNextEdge(edges []graph.Edge[PheromoneDistanceAngle, vector.Coordinate]) graph.Edge[PheromoneDistanceAngle, vector.Coordinate] {
	q := rand.Float64()

	if q > a.cfg.RandomChance {
		return a.chooseNextEdgeWeighted(edges)
	} else {
		n := rand.Intn(len(edges))
		return edges[n]
	}
}

func (a *Ant) chooseNextEdgeWeighted(edges []graph.Edge[PheromoneDistanceAngle, vector.Coordinate]) graph.Edge[PheromoneDistanceAngle, vector.Coordinate] {
	weights := make([]float64, len(edges))
	for i, e := range edges {
		weights[i] = e.Weight.Pheromone * a.cfg.PheromoneWeight / (e.Weight.DistanceAngle.Distance * a.cfg.DistanceWeight)
	}
	n := pkg.ChooseWeighted(weights)
	return edges[n]
}

func updatePheromone(cfg Config, totalEdgeLength float64, g *graph.WeightedGraph[vector.Coordinate, PheromoneDistanceAngle], shortestPath []graph.Edge[PheromoneDistanceAngle, vector.Coordinate]) {
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
		newPheromone := e.Weight.Pheromone * totalEdgeLength / totalDistance
		g.UpdateEdge(e.From, e.To, PheromoneDistanceAngle{
			Pheromone:     newPheromone,
			DistanceAngle: e.Weight.DistanceAngle,
		})
	}
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
