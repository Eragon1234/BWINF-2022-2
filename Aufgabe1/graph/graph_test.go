package graph

import (
	"testing"
)

func TestWeightedGraph_AddVertex(t *testing.T) {
	graph := NewWeightedGraph[int, int]()

	graph.AddVertex("A", 1)
	graph.AddVertex("B", 2)
	graph.AddVertex("C", 3)

	if graph.Vertices["A"].Value != 1 {
		t.Error("Wrong value for A")
	}

	if graph.Vertices["B"].Value != 2 {
		t.Error("Wrong value for B")
	}

	if graph.Vertices["C"].Value != 3 {
		t.Error("Wrong value for C")
	}
}

func TestWeightedGraph_AddEdge(t *testing.T) {
	graph := NewWeightedGraph[int, int]()

	graph.AddVertex("A", 1)
	graph.AddVertex("B", 2)
	graph.AddVertex("C", 3)

	graph.AddEdge(graph.Vertices["A"], graph.Vertices["B"], 3)

	if !graph.GetEdge(graph.Vertices["A"], graph.Vertices["B"]).Exists {
		t.Error("No edge between A and B")
	}

	if graph.GetEdge(graph.Vertices["B"], graph.Vertices["A"]).Exists {
		t.Error("Edge between B and A exists")
	}

	if graph.GetEdge(graph.Vertices["A"], graph.Vertices["B"]).Weight != 3 {
		t.Error("Wrong weight between A and B")
	}

	if graph.GetEdge(graph.Vertices["A"], graph.Vertices["C"]).Exists {
		t.Error("Non existing edge between A and C exists")
	}

	if graph.GetEdge(graph.Vertices["C"], graph.Vertices["A"]).Exists {
		t.Error("Non existing edge between C and A exists")
	}

	if graph.GetEdge(graph.Vertices["B"], graph.Vertices["C"]).Exists {
		t.Error("Non existing edge between B and C exists")
	}

	if graph.GetEdge(graph.Vertices["C"], graph.Vertices["B"]).Exists {
		t.Error("Non existing edge between C and B exists")
	}
}
