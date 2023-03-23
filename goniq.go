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

func Add[T Ordered](slice []T, s T) []T {
	if len(slice) == 0 {
		return []T{s}
	}

	slice = sortSlice(slice)
	idx := sort.Search(len(slice), func(i int) bool { return slice[i] >= s })
	if idx == len(slice) || slice[idx] != s {
		return append(slice[:idx], append([]T{s}, slice[idx:]...)...)
	}
	return slice
}

func Remove[T Ordered](slice []T, s T) []T {
	if len(slice) == 0 {
		return slice
	}

	slice = sortSlice(slice)
	idx := sort.Search(len(slice), func(i int) bool { return slice[i] >= s })
	if idx < len(slice) && slice[idx] == s {
		if idx == len(slice)-1 {
			return slice[:idx]
		}
		copy(slice[idx:], slice[idx+1:])
		return slice[:len(slice)-1]
	}
	return slice
}
