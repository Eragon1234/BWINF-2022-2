package graph

type WeightedGraph[Data any, Weight any] struct {
	Vertices        map[string]*Vertex[Data]
	adjacencyMatrix [][]Edge[Weight, Data]
}

func NewWeightedGraph[Data any, Weight any]() *WeightedGraph[Data, Weight] {
	graph := &WeightedGraph[Data, Weight]{
		Vertices:        make(map[string]*Vertex[Data]),
		adjacencyMatrix: [][]Edge[Weight, Data]{},
	}

	return graph
}

func (g *WeightedGraph[D, W]) AddVertex(name string, value D) {
	g.Vertices[name] = &Vertex[D]{Name: name, Value: value, Index: len(g.Vertices)}
	g.adjacencyMatrix = append(g.adjacencyMatrix, make([]Edge[W, D], len(g.Vertices)-1))
	for i := range g.adjacencyMatrix {
		g.adjacencyMatrix[i] = append(g.adjacencyMatrix[i], Edge[W, D]{
			From: &Vertex[D]{},
			To:   &Vertex[D]{},
		})
	}
}

func (g *WeightedGraph[D, W]) AddEdge(from, to *Vertex[D], weight W) {
	g.adjacencyMatrix[from.Index][to.Index] = Edge[W, D]{From: from, To: to, Exists: true, Weight: weight}
	g.adjacencyMatrix[to.Index][from.Index] = Edge[W, D]{From: to, To: from, Exists: true, Weight: weight}
}

func (g *WeightedGraph[D, W]) GetEdge(from, to *Vertex[D]) Edge[W, D] {
	return g.adjacencyMatrix[from.Index][to.Index]
}
