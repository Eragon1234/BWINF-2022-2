package graph

type Edge[T any] struct {
	Exists   bool
	From, To int
	Weight   T
}
