package utils

type Set struct {
	m map[string]bool
}

func (s *Set) Add(key string) {
	s.m[key] = true
}

func (s *Set) Contains(key string) bool {
	_, ok := s.m[key]
	return ok
}

func (s *Set) Remove(key string) {
	delete(s.m, key)
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) Clear() {
	s.m = make(map[string]bool)
}

func (s *Set) IsEmpty() bool {
	return s.Size() == 0
}
