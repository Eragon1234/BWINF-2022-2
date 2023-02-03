package graph

type WeightedGraph[M any, T any] struct {
	adjacencyMatrix [][]Edge[T]
	Vertices        []Vertex[M]
}

func NewWeightedGraph[M any, T any](size int) *WeightedGraph[M, T] {
	graph := &WeightedGraph[M, T]{
		Vertices:        make([]Vertex[M], 0, size),
		adjacencyMatrix: make([][]Edge[T], size),
	}

	for i := range graph.adjacencyMatrix {
		graph.adjacencyMatrix[i] = make([]Edge[T], size)
	}

	return graph
}

func (g *WeightedGraph[M, T]) AddVertex(value M) {
	g.Vertices = append(g.Vertices, Vertex[M]{Index: len(g.Vertices), Value: value})
	g.adjacencyMatrix = append(g.adjacencyMatrix, make([]Edge[T], len(g.Vertices)))
	for i := range g.adjacencyMatrix {
		g.adjacencyMatrix[i] = append(g.adjacencyMatrix[i], Edge[T]{})
	}
}

func (g *WeightedGraph[M, T]) AddEdge(from, to int, weight T) {
	g.adjacencyMatrix[from][to] = Edge[T]{From: from, To: to, Weight: weight}
	g.adjacencyMatrix[to][from] = Edge[T]{From: to, To: from, Weight: weight}
}

func (g *WeightedGraph[M, T]) GetEdge(from, to int) Edge[T] {
	return g.adjacencyMatrix[from][to]
}
