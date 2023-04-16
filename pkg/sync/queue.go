package sync

import (
	"BWINF/pkg/queue"
	"sync"
)

// PriorityQueue is a wrapper around queue.PriorityQueue that allows safe use by multiple goroutines without additional locking or coordination.
type PriorityQueue[T any] struct {
	mu sync.Mutex
	pg queue.PriorityQueue[T]
}

func (pq *PriorityQueue[T]) Push(val T, priority int) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	pq.pg.Push(val, priority)
}

func (pq *PriorityQueue[T]) Pop() (T, bool) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	return pq.pg.Pop()
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.pg.Len()
}
