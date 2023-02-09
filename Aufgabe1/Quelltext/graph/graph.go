package graph

type WeightedGraph[Data any, Weight any] struct {
	Vertices        map[string]Vertex[Data]
	adjacencyMatrix [][]Edge[Weight]
}

func NewWeightedGraph[Data any, Weight any](size int) *WeightedGraph[Data, Weight] {
	graph := &WeightedGraph[Data, Weight]{
		Vertices:        make(map[string]Vertex[Data], size),
		adjacencyMatrix: make([][]Edge[Weight], size),
	}

	for i := range graph.adjacencyMatrix {
		graph.adjacencyMatrix[i] = make([]Edge[Weight], size)
	}

	return graph
}

func (g *WeightedGraph[D, W]) AddVertex(name string, value D) {
	g.Vertices[name] = Vertex[D]{Name: name, Value: value}
	g.adjacencyMatrix = append(g.adjacencyMatrix, make([]Edge[W], len(g.Vertices)))
	for i := range g.adjacencyMatrix {
		g.adjacencyMatrix[i] = append(g.adjacencyMatrix[i], Edge[W]{})
	}
}

func (g *WeightedGraph[D, W]) AddEdge(from, to int, weight W) {
	g.adjacencyMatrix[from][to] = Edge[W]{From: from, To: to, Exists: true, Weight: weight}
	g.adjacencyMatrix[to][from] = Edge[W]{From: to, To: from, Exists: true, Weight: weight}
}

func (g *WeightedGraph[D, W]) GetEdge(from, to int) Edge[W] {
	return g.adjacencyMatrix[from][to]
}
