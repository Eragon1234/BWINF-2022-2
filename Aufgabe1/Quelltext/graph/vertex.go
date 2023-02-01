package graph

type Vertex[M any] struct {
	Index int
	Value M
}
