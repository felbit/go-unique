package goniq

import "sort"

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Int interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Ordered interface {
	Float | Int | ~string
}

func sortSlice[T Ordered](s []T) []T {
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	return s
}

// NewSet takes a slice of a `comparable` type and returns possibly reduced set of unique entries of the original slice.
// Set is unsorted. The order of elements will be the first unique appearance of elements in the original slice.
//
// Example:
//
//	NewSet([]int{3, 1, 5, 2, 3, 2, 1}) // => []int{3, 1, 5, 2}
func NewSet[C comparable](slice []C) []C {
	if len(slice) < 2 {
		return slice
	}

	unique := make(map[C]bool)
	result := make([]C, 0, len(slice))
	for _, e := range slice {
		if !unique[e] {
			unique[e] = true
			result = append(result, e)
		}
	}
	return result
}

func Add[T Ordered](slice []T, element T) []T {
	if len(slice) == 0 {
		return []T{element}
	}

	slice = sortSlice(slice)
	idx := sort.Search(len(slice), func(i int) bool { return slice[i] >= element })
	if idx == len(slice) || slice[idx] != element {
		return append(slice[:idx], append([]T{element}, slice[idx:]...)...)
	}
	return slice
}

func Remove[T Ordered](slice []T, element T) []T {
	if len(slice) == 0 {
		return slice
	}

	slice = sortSlice(slice)
	idx := sort.Search(len(slice), func(i int) bool { return slice[i] >= element })
	if idx < len(slice) && slice[idx] == element {
		if idx == len(slice)-1 {
			return slice[:idx]
		}
		copy(slice[idx:], slice[idx+1:])
		return slice[:len(slice)-1]
	}
	return slice
}
