package graph

type Edge[T any, M any] struct {
	Exists   bool
	From, To *Vertex[M]
	Weight   T
}
