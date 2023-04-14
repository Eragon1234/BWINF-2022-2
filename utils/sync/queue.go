package sync

import (
	"BWINF/utils/queue"
	"sync"
)

// PriorityQueue is a wrapper around queue.PriorityQueue that allows safe use by multiple goroutines without additional locking or coordination.
type PriorityQueue[T any] struct {
	mu sync.Mutex
	pg queue.PriorityQueue[T]
}

func NewPriorityQueue[T any]() *PriorityQueue[T] {
	return &PriorityQueue[T]{
		pg: queue.PriorityQueue[T]{},
	}
}

func (pq *PriorityQueue[T]) Push(val queue.Item[T]) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	pq.pg.Push(val)
}

func (pq *PriorityQueue[T]) Pop() (queue.Item[T], bool) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	return pq.pg.Pop()
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.pg.Len()
}
