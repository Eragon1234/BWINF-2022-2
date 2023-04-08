package utils

import "container/heap"

// An PQItem is something we manage in a Priority queue.
type PQItem[T any] struct {
	Value    T
	Priority uint8
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A innerPriorityQueue implements heap.Interface and holds Items.
type innerPriorityQueue[T any] []PQItem[T]

func (pq innerPriorityQueue[T]) Len() int { return len(pq) }

func (pq innerPriorityQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, Priority, so we use greater than here.
	return pq[i].Priority > pq[j].Priority
}

func (pq innerPriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *innerPriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(PQItem[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *innerPriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type PriorityQueue[T any] struct {
	ipq innerPriorityQueue[T]
}

func (pq *PriorityQueue[T]) Push(x PQItem[T]) {
	heap.Push(&pq.ipq, x)
}

func (pq *PriorityQueue[T]) Pop() (PQItem[T], bool) {
	if pq.ipq.Len() == 0 {
		return PQItem[T]{}, false
	}
	return heap.Pop(&pq.ipq).(PQItem[T]), true
}

func (pq *PriorityQueue[T]) Update(item PQItem[T], value T, priority uint8) {
	item.Value = value
	item.Priority = priority
	heap.Fix(&pq.ipq, item.index)
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.ipq.Len()
}
