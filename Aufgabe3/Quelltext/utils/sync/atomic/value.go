package atomic

import "sync/atomic"

// Value is a generic wrapper around atomic.value that allows type safety
type Value[T any] struct {
	v atomic.Value
}

func (v *Value[T]) Load() (T, bool) {
	value, ok := v.v.Load().(T)
	return value, ok
}

func (v *Value[T]) Store(value T) {
	v.v.Store(value)
}

func (v *Value[T]) CompareAndSwap(old, n T) bool {
	_, ok := v.Load()
	if !ok {
		// if the value is not set, set it
		v.Store(n)
		return true
	}
	return v.v.CompareAndSwap(old, n)
}

func (v *Value[T]) Swap(n T) (T, bool) {
	old, ok := v.v.Swap(n).(T)
	return old, ok
}
