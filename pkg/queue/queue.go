package queue

import "container/heap"

// An Item is something we manage in a Priority queue.
type Item[T any] struct {
	Value    T
	Priority int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type PriorityQueue[T any] struct {
	ipq internalPriorityQueue[T]
}

func (pq *PriorityQueue[T]) Push(x T, priority int) {
	item := Item[T]{
		Value:    x,
		Priority: priority,
		index:    pq.ipq.Len(),
	}
	heap.Push(&pq.ipq, item)
}

func (pq *PriorityQueue[T]) Pop() (T, bool) {
	if pq.ipq.Len() == 0 {
		return *new(T), false
	}
	return heap.Pop(&pq.ipq).(Item[T]).Value, true
}

func (pq *PriorityQueue[T]) Update(item Item[T], value T, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(&pq.ipq, item.index)
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.ipq.Len()
}
