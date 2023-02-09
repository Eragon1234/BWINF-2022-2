package graph

import (
	"Aufgabe1/graph/vector"
	"bufio"
	"io"
)

// ParseComplete parses a complete weighted graph from a reader with the weight as DistanceAngle
func ParseComplete(reader io.Reader) (WeightedGraph[vector.Coordinate, DistanceAngle], error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var coordinates []vector.Coordinate

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		c, err := vector.ParseCoordinate(line)
		if err != nil {
			return WeightedGraph[vector.Coordinate, DistanceAngle]{}, err
		}
		coordinates = append(coordinates, c)
	}

	weightedGraph := *NewWeightedGraph[vector.Coordinate, DistanceAngle](len(coordinates))
	for _, c := range coordinates {
		weightedGraph.AddVertex(c.String(), c)
	}

	for i, vertex := range weightedGraph.Vertices {
		for j, otherVertex := range weightedGraph.Vertices {
			if i == j {
				continue
			}
			weightedGraph.AddEdge(vertex.Index, otherVertex.Index, DistanceAngle{
				Distance: vector.Distance(vertex.Value, otherVertex.Value),
				Angle:    vector.Angle(vertex.Value, otherVertex.Value),
			})
		}
	}

	return weightedGraph, nil
}

type DistanceAngle struct {
	Distance float64
	Angle    float64
}
