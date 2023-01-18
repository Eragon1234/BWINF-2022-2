package utils

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type TicketSystem[T any] struct {
	data  []T
	mutex sync.RWMutex
	count atomic.Uint64
	done  atomic.Uint64
}

func (ts *TicketSystem[T]) Put(value T) {
	t := ts.Ticket()
	ts.Swap(t, value)
	ts.Done(t)
}

func (ts *TicketSystem[T]) Ticket() uint64 {
	ts.mutex.Lock()
	var newT T
	ts.data = append(ts.data, newT)
	ts.mutex.Unlock()
	return ts.count.Add(1) - 1
}

func (ts *TicketSystem[T]) Swap(ticket uint64, value T) {
	ts.mutex.RLock()
	ts.data[ticket] = value
	ts.mutex.RUnlock()
}

func (ts *TicketSystem[T]) Done(ticket uint64) {
	for !(ts.done.CompareAndSwap(ticket, ticket+1)) {
		runtime.Gosched()
	}
}

func (ts *TicketSystem[T]) GetDone() []T {
	return ts.data[:ts.done.Load()]
}
