package graph

import (
	"BWINF/Aufgabe1/vector"
	"BWINF/utils"
	"BWINF/utils/slice"
)

func VisitAllShortestEdge(g WeightedGraph[vector.Coordinate, DistanceAngle]) []Edge[DistanceAngle, vector.Coordinate] {
	paths := make([][]Edge[DistanceAngle, vector.Coordinate], 0, len(g.Vertices))
	for _, v := range g.Vertices {
		paths = append(paths, VisitAllShortestEdgeWithStart(g, v))
	}
	return slice.MinFunc(paths, func(a, b []Edge[DistanceAngle, vector.Coordinate]) bool {
		return LengthOfPath(a) < LengthOfPath(b)
	})
}

func VisitAllShortestEdgeWithStart(g WeightedGraph[vector.Coordinate, DistanceAngle], start *Vertex[vector.Coordinate]) []Edge[DistanceAngle, vector.Coordinate] {
	var visited utils.Set[*Vertex[vector.Coordinate]]
	visited.Add(start)
	var path []Edge[DistanceAngle, vector.Coordinate]

	for visited.Size() < len(g.Vertices) {
		var curr Edge[DistanceAngle, vector.Coordinate]
		var edges []Edge[DistanceAngle, vector.Coordinate]
		if len(path) > 0 {
			curr = path[len(path)-1]
			edges = slice.FilterFunc(g.GetEdges(curr.To), func(e Edge[DistanceAngle, vector.Coordinate]) bool {
				return visited.Contains(e.To) || vector.TurnAngle(curr.Weight.Angle, e.Weight.Angle) > 90
			})
		} else {
			edges = slice.FilterFunc(g.GetEdges(start), func(e Edge[DistanceAngle, vector.Coordinate]) bool {
				return visited.Contains(e.To)
			})
		}
		if len(edges) == 0 {
			// we could backtrack,
			// but in some cases we need to accept a turn angle greater than 90 because we mostly can't have a path with all turn angles less than 90
			// figuring out when to accept a turn angle greater than 90 takes too much time,
			// so we just accept a turn angle greater than 90
			edges = slice.FilterFunc(g.GetEdges(curr.To), func(e Edge[DistanceAngle, vector.Coordinate]) bool {
				return visited.Contains(e.To)
			})
		}
		minEdge := slice.MinFunc(edges, func(a, b Edge[DistanceAngle, vector.Coordinate]) bool {
			return a.Weight.Distance < b.Weight.Distance
		})
		visited.Add(minEdge.To)
		path = append(path, minEdge)
	}
	return path
}

func LengthOfPath(path []Edge[DistanceAngle, vector.Coordinate]) float64 {
	return slice.SumFunc(path, func(a float64, b Edge[DistanceAngle, vector.Coordinate]) float64 {
		return a + b.Weight.Distance
	})
}
