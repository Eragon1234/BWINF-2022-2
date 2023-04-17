package ant

import (
	"BWINF/Aufgabe1/graph"
	"BWINF/Aufgabe1/vector"
)

func TransformGraph(g graph.WeightedGraph[vector.Coordinate, graph.DistanceAngle]) *graph.WeightedGraph[vector.Coordinate, PheromoneDistanceAngle] {
	pheromoneGraph := graph.NewWeightedGraph[vector.Coordinate, PheromoneDistanceAngle]()
	for _, v := range g.Vertices {
		pheromoneGraph.AddVertex(v.Name, v.Value)
	}
	for _, v := range g.Vertices {
		for _, e := range g.GetEdges(v) {
			pheromoneGraph.AddEdge(pheromoneGraph.Vertices[e.From.Name], pheromoneGraph.Vertices[e.To.Name], PheromoneDistanceAngle{
				Pheromone:     DefaultConfig.PheromoneAmount,
				DistanceAngle: e.Weight,
			})
		}
	}
	return pheromoneGraph
}

func transformEdges(edges []graph.Edge[graph.DistanceAngle, vector.Coordinate]) []graph.Edge[PheromoneDistanceAngle, vector.Coordinate] {
	pheromoneEdges := make([]graph.Edge[PheromoneDistanceAngle, vector.Coordinate], len(edges))
	for i, e := range edges {
		pheromoneEdges[i] = graph.Edge[PheromoneDistanceAngle, vector.Coordinate]{
			From:   e.From,
			To:     e.To,
			Weight: PheromoneDistanceAngle{Pheromone: DefaultConfig.PheromoneAmount, DistanceAngle: e.Weight},
		}
	}
	return pheromoneEdges
}
