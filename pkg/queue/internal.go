package queue

// internalPriorityQueue implements heap.Interface and holds Items.
// it returns items with smaller priority first
type internalPriorityQueue[T any] []Item[T]

func (pq internalPriorityQueue[T]) Len() int { return len(pq) }

func (pq internalPriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq internalPriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *internalPriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(Item[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *internalPriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
