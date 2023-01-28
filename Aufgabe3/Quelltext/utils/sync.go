package utils

import (
	"sync"
	"sync/atomic"
)

// AtomicValue is a generic wrapper around atomic.value that allows type safety
type AtomicValue[T any] struct {
	v atomic.Value
}

func (v *AtomicValue[T]) Load() (T, bool) {
	value, ok := v.v.Load().(T)
	return value, ok
}

func (v *AtomicValue[T]) Store(value T) {
	v.v.Store(value)
}

func (v *AtomicValue[T]) CompareAndSwap(old, new T) bool {
	return v.v.CompareAndSwap(old, new)
}

func (v *AtomicValue[T]) Swap(new T) (T, bool) {
	old, ok := v.v.Swap(new).(T)
	return old, ok
}

// SyncMap is a generic wrapper around sync.Map that allows type safety
type SyncMap[K comparable, V any] struct {
	m sync.Map
}

func (m *SyncMap[K, V]) Delete(key K) {
	m.m.Delete(key)
}

func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.m.Load(key)
	if !ok {
		return value, ok
	}
	return v.(V), ok
}

func (m *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := m.m.LoadAndDelete(key)
	if !loaded {
		return value, loaded
	}
	return v.(V), loaded
}

func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	a, loaded := m.m.LoadOrStore(key, value)
	return a.(V), loaded
}

func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func (m *SyncMap[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}
