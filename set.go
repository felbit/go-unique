package goniq

import "sync"

// Set is a set-like data structure that holds elements of `Ordered` type.
// It implements goniq's functions in a thread-safe manner.
type Set[O Ordered] struct {
	sync.RWMutex
	set *[]O
}

// NewSet returns a new Set (a slice of the given type of Ordered with zero elements).
func NewSet[O Ordered]() *Set[O] {
	return &Set[O]{set: &[]O{}}
}

// Contains returns true if the given element is present in the set.
func (s *Set[O]) Contains(element O) bool {
	s.RLock()
	defer s.RUnlock()
	return Contains(s.set, element)
}

// Add inserts an element of `Ordered` type into the set. If the element is already present in the slice,
// the set is left unchanged (will **not** mutate the set). Otherwise, the element is added at the end of the set.
//
// Example usage:
//
//	set := NewSet[int]() // set == []
//	set.Add(1) //  => [1]
//	set.Add(1) //  => [1]
//	set.Add(2) //  => [1, 2]
func (s *Set[O]) Add(element O) {
	s.Lock()
	defer s.Unlock()
	Add(s.set, element)
}

// Append inserts multiple elements of `Ordered` type into the set. The function modifies the set in place and adds new
// unique elements at the end. Will ignore elements that are already present.
//
// Example:
//
//	set := NewSet[int]()   // set == []
//	set.Add(1)             //  => [1]
//	set.Append(2, 3, 1, 4) //  => [1, 2, 3, 4]
func (s *Set[O]) Append(elements ...O) {
	s.Lock()
	defer s.Unlock()
	Append(s.set, elements...)
}

// Remove removes the given element from the set. The function modifies the set in place.
//
// Example:
//
//	set := NewSet[int]() // set == []
//	set.Add(1) //  => [1]
//	set.Add(2) //  => [1, 2]
//	set.Add(3) //  => [1, 2, 3]
//	set.Remove(2) //  => [1, 3]
func (s *Set[O]) Remove(element O) {
	s.Lock()
	defer s.Unlock()
	Remove(s.set, element)
}

// Size returns the number of elements in the set.
func (s *Set[O]) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(*s.set)
}

// RemoveAll removes all elements from the set (it replaces the inner slice with a new empty slice).
func (s *Set[O]) RemoveAll() {
	s.Lock()
	defer s.Unlock()
	*s.set = []O{}
}
