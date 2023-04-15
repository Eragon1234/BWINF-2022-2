package set

type Set[T comparable] struct {
	m map[T]bool
}

func New[T comparable](capacity int) *Set[T] {
	return &Set[T]{m: make(map[T]bool, capacity)}
}

func (s *Set[T]) Add(key T) {
	if s.m == nil {
		s.m = make(map[T]bool)
	}
	s.m[key] = true
}

func (s *Set[T]) Contains(key T) bool {
	_, ok := s.m[key]
	return ok
}

func (s *Set[T]) Remove(key T) {
	delete(s.m, key)
}

func (s *Set[T]) Size() int {
	return len(s.m)
}

func (s *Set[T]) Clear() {
	s.m = make(map[T]bool)
}

func (s *Set[T]) IsEmpty() bool {
	return s.Size() == 0
}
