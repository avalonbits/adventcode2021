// Package collecetions implements multiple generic collections.
package collections

// Set is a collection of unordered unique values.
type Set[T comparable] struct {
	set map[T]struct{}
}

// NewSet returns an initialized set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		set: map[T]struct{}{},
	}
}

// Add adds an element to the set and returns itself.
func (s *Set[T]) Add(v T) *Set[T] {
	s.set[v] = struct{}{}
	return s
}

// Insert adds multiple elements to the set and returns itself.
func (s *Set[T]) Insert(vs ...T) *Set[T] {
	for _, v := range vs {
		s.Add(v)
	}
	return s
}

// Size returns the number of elements in the set
func (s *Set[t]) Size() int {
	return len(s.set)
}

// In returns true if v is in the set.
func (s *Set[T]) In(v T) bool {
	_, ok := s.set[v]
	return ok
}

// Values returns a slice with all elements, in no particular order.
func (s *Set[T]) Values() []T {
	values := make([]T, len(s.set))
	for key := range s.set {
		values = append(values, key)
	}
	return values
}
