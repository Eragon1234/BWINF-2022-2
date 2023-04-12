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
	for i := range g.adjacencyMatrix[len(g.adjacencyMatrix)-1] {
		g.adjacencyMatrix[len(g.adjacencyMatrix)-1][i] = Edge[W, D]{
			From: g.Vertices[name],
			To:   g.Vertices[name],
		}
	}
	for i := range g.adjacencyMatrix {
		g.adjacencyMatrix[i] = append(g.adjacencyMatrix[i], Edge[W, D]{
			From: g.Vertices[name],
			To:   g.Vertices[name],
		})
	}
}

func (g *WeightedGraph[D, W]) AddEdge(from, to *Vertex[D], weight W) {
	g.adjacencyMatrix[from.Index][to.Index] = Edge[W, D]{From: from, To: to, Exists: true, Weight: weight}
}

func (g *WeightedGraph[D, W]) GetEdge(from, to *Vertex[D]) Edge[W, D] {
	return g.adjacencyMatrix[from.Index][to.Index]
}

func (g *WeightedGraph[D, W]) UpdateEdge(from, to *Vertex[D], weight W) {
	g.adjacencyMatrix[from.Index][to.Index].Weight = weight
}

func (g WeightedGraph[Data, Weight]) GetEdges(v *Vertex[Data]) []Edge[Weight, Data] {
	return g.adjacencyMatrix[v.Index]
}

func (g *WeightedGraph[D, W]) Copy() WeightedGraph[D, W] {
	newGraph := NewWeightedGraph[D, W]()
	for _, v := range g.Vertices {
		newGraph.AddVertex(v.Name, v.Value)
	}
	for _, v := range g.Vertices {
		for _, e := range g.GetEdges(v) {
			newGraph.AddEdge(newGraph.Vertices[e.From.Name], newGraph.Vertices[e.To.Name], e.Weight)
		}
	}
	return *newGraph
}
