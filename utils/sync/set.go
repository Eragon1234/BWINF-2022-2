package sync

type Set[T comparable] struct {
	m Map[T, bool]
}

func (s *Set[T]) Add(key T) {
	s.m.Store(key, true)
}

func (s *Set[T]) Contains(key T) bool {
	_, ok := s.m.Load(key)
	return ok
}

func (s *Set[T]) Remove(key T) {
	s.m.Delete(key)
}

func (s *Set[T]) Clear() {
	s.m = Map[T, bool]{}
}
