package sync

import (
	"BWINF/utils"
	"BWINF/utils/queue"
)

type Representer interface {
	Represent() string
}

type PriorityQueue[T Representer] struct {
	lock    chan struct{}
	pg      queue.PriorityQueue[T]
	visited utils.Set[string]
}

func NewPriorityQueue[T Representer]() *PriorityQueue[T] {
	lock := make(chan struct{}, 1)
	lock <- struct{}{}
	return &PriorityQueue[T]{
		lock:    lock,
		pg:      queue.PriorityQueue[T]{},
		visited: utils.Set[string]{},
	}
}

func (pq *PriorityQueue[T]) Push(val queue.Item[T]) {
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

func (pq *PriorityQueue[T]) Pop() (queue.Item[T], bool) {
	<-pq.lock
	defer func() {
		pq.lock <- struct{}{}
	}()
	return pq.pg.Pop()
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.pg.Len()
}
