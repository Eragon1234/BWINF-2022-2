package graph

type WeightedGraph[M any, T any] struct {
	adjacencyMatrix [][]Edge[T]
	Vertices        []Vertex[M]
}

func NewWeightedGraph[M any, T any](size int) *WeightedGraph[M, T] {
	graph := &WeightedGraph[M, T]{
		adjacencyMatrix: make([][]Edge[T], size),
	}

	for i := range graph.adjacencyMatrix {
		graph.adjacencyMatrix[i] = make([]Edge[T], size)
	}

	return graph
}
