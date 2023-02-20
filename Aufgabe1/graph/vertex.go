package graph

type Vertex[M any] struct {
	Name  string
	Index int
	Value M
}
