package queue

import "container/heap"

// An Item is something we manage in a Priority queue.
type Item[T any] struct {
	Value    T
	Priority uint8
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type PriorityQueue[T any] struct {
	ipq internalPriorityQueue[T]
}

func (pq *PriorityQueue[T]) Push(x Item[T]) {
	heap.Push(&pq.ipq, x)
}

func (pq *PriorityQueue[T]) Pop() (Item[T], bool) {
	if pq.ipq.Len() == 0 {
		return Item[T]{}, false
	}
	return heap.Pop(&pq.ipq).(Item[T]), true
}

func (pq *PriorityQueue[T]) Update(item Item[T], value T, priority uint8) {
	item.Value = value
	item.Priority = priority
	heap.Fix(&pq.ipq, item.index)
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.ipq.Len()
}
