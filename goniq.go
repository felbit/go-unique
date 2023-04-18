package goniq

import (
	"regexp"
	"sort"
)

func sortSlice[T Ordered](s []T) []T {
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	return s
}

// Contains inserts an element of type T into a sorted slice of elements of type T.
// The slice is mutated in place.
//
// Example usage:
//
//	slice := []int{1, 2, 3}
//	Contains(&slice, 3) // => true
//	Contains(&slice, 4) // => false
func Contains[T Ordered](s *[]T, e T) bool {
	idx := sort.Search(len(*s), func(i int) bool { return (*s)[i] >= e })
	return idx < len(*s) && (*s)[idx] == e
}

// RemoveDuplicates takes a slice of a `comparable` type and mutates it to only contain unique entries.
// The result keeps the order of first occurrence of elements.
//
// Example usage:
//
//	slice := []int{3, 1, 5, 2, 3, 2, 1}
//	RemoveDuplicates(&slice) // => []int{3, 1, 5, 2}
func RemoveDuplicates[C comparable](slice *[]C) {
	if len(*slice) < 2 {
		return
	}

	seen := make(map[C]bool)
	idx := 0
	for _, e := range *slice {
		if !seen[e] {
			seen[e] = true
			(*slice)[idx] = e
			idx++
		}
	}
	*slice = (*slice)[:idx]
}

// Add inserts an element of type T into a sorted slice of elements of type T.
// If the element is already present in the slice, the original slice is returned unchanged.
// Otherwise, a new slice is returned with the element inserted at the end.
//
// Example usage:
//
//	slice := []int{1, 2, 4, 5}
//	Add(&slice, 3) //  => []int{1, 2, 4, 5, 3}
func Add[T Ordered](slice *[]T, element T) {
	// check if the element is in the slice
	sorted := make([]T, len(*slice))
	copy(sorted, *slice)

	sorted = sortSlice(sorted)
	idx := sort.Search(len(sorted), func(i int) bool { return sorted[i] >= element })

	if idx == len(sorted) || sorted[idx] != element {
		*slice = append(*slice, element)
		return
	}
}

// Append inserts multiple elements of type T into a sorted slice of elements of type T.
// The function modifies the slice in place and adds new unique elements at the end.
// Will ignore elements that are already present.
//
// Example:
//
//	slice := []int{1, 2, 4, 5}
//	Append(&slice, 3, 6, 7, 4, 5) // => []int{1, 2, 4, 5, 3, 6, 7}
func Append[T Ordered](slice *[]T, elements ...T) {
	for _, e := range elements {
		Add(slice, e)
	}
}

// Remove removes all occurrences of the given element from the given slice of elements of type T.
// The function modifies the slice in place and returns it **sorted**.
//
// Example:
//
//	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
//	Remove(&slice, 5) // removes all 5's from the slice
//	fmt.Println(nums) // prints [1, 1, 2, 3, 3, 4, 6, 9]
func Remove[T Ordered](slice *[]T, element T) {
	if len(*slice) == 0 {
		return
	}

	*slice = sortSlice(*slice)
	for {
		idx := sort.Search(len(*slice), func(i int) bool { return (*slice)[i] >= element })
		if idx >= len(*slice) || (*slice)[idx] != element {
			break
		}
		if idx == len(*slice)-1 {
			*slice = (*slice)[:idx]
		} else {
			copy((*slice)[idx:], (*slice)[idx+1:])
			*slice = (*slice)[:len(*slice)-1]
		}
	}
}

// RemoveStringsAkin removes strings from a slice of strings that match partially.
// The function mutates the slice in place and is is case-sensitive.
//
// Example:
//
//	slice := []string{"foobar", "foobaz", "fobar", "quuz", "qufooz"}
//	RemoveStringAkin(&slice, "foo") // => []string{"fobar", "quuz"}
func RemoveStringsAkin(slice *[]string, str string) {
	re := regexp.MustCompile(str)
	idx := 0

	for _, s := range *slice {
		if !re.MatchString(s) {
			(*slice)[idx] = s
			idx++
		}
	}

	*slice = (*slice)[:idx]
}
