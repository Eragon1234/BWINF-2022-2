package sync

import (
	"BWINF/utils"
)

type Representer interface {
	Represent() string
}

type PriorityQueue[T Representer] struct {
	lock    chan struct{}
	pg      utils.PriorityQueue[T]
	visited utils.Set[string]
}

func NewPriorityQueue[T Representer]() *PriorityQueue[T] {
	lock := make(chan struct{}, 1)
	lock <- struct{}{}
	return &PriorityQueue[T]{
		lock:    lock,
		pg:      utils.PriorityQueue[T]{},
		visited: utils.Set[string]{},
	}
}

func (pq *PriorityQueue[T]) Push(val utils.PQItem[T]) {
	<-pq.lock
	defer func() {
		pq.lock <- struct{}{}
	}()
	representation := val.Value.Represent()
	if pq.visited.Contains(representation) {
		return
	}
	pq.visited.Add(representation)
	pq.pg.Push(val)
}

func (pq *PriorityQueue[T]) Pop() (utils.PQItem[T], bool) {
	<-pq.lock
	defer func() {
		pq.lock <- struct{}{}
	}()
	return pq.pg.Pop()
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.pg.Len()
}
