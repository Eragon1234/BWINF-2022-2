package graph

type Vertex[M any] struct {
	Name  string
	Value M
}
