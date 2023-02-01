package graph

type Edge[T any] struct {
	From, To int
	Weight   T
}
