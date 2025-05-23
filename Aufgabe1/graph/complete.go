package graph

import (
	"BWINF/Aufgabe1/vector"
	"bufio"
	"io"
)

// ParseComplete parses a complete weighted graph with the weight as DistanceAngle
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

	weightedGraph := *NewWeightedGraph[vector.Coordinate, DistanceAngle]()
	for _, c := range coordinates {
		// because the coordinates are unique, we can use them as vertex names since we don't have names
		weightedGraph.AddVertex(c.String(), c)
	}

	for i, vertex := range weightedGraph.Vertices {
		for j, otherVertex := range weightedGraph.Vertices {
			if i == j {
				continue
			}
			weightedGraph.AddEdge(vertex, otherVertex, DistanceAngle{
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
