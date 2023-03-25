# goniq
![Go Version](https://img.shields.io/badge/Go-v1.20-blue)
![Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)

*go unique* ... simple set-like operations for Go slices.

>   :warning: This library is not considered stable. Therefore, caution is advised using it for production for now.

## Goals

- "out of the way"
- generic
- idempotent

## Functions

### RemoveDuplicates

```go
func RemoveDuplicates[C comparable](*[]C)
```

This function takes a slice of a `comparable` type and returns possibly reduced set of unique entries of the original slice.
Set is unsorted. The order of elements will be the first unique appearance of elements in the original slice.

### Add

```go
func Add[T Ordered](*[]T, T)
```

This function takes a slice of T and an element of T as input,
and returns a new sorted slice of T with the input element added to it.
If the input slice is empty, it creates a new slice with only the input string.
If the input string is already in the slice, it returns the original slice, but sorted.


### Remove

```go
func Remove[T Ordered](*[]T, T)
```

>   :warning: `Remove` currently returns a **sorted slice**, even if the elements fo not change. This might change in the future.

This function takes a slice of T and an element of T as input,
and returns a new sorted slice of T with the input element removed from it.
If the input slice is empty or the input element is not in the slice,
it returns the original slice sorted.

## Usage

```go
package main

import (
	"fmt"
	"github.com/felbit/goniq"
)

func main() {
	s := []string{"apple", "banana", "apple", "cherry"}
	goniq.RemoveDuplicates(&s) // returns ["apple", "banana", "cherry"]
	goniq.Add(&s, "banana") // returns ["apple", "banana", "cherry"]
	goniq.Add(&s, "date") // returns ["apple", "banana", "cherry", "date"]
	goniq.Remove(&s, "banana") // returns ["apple", "cherry", "date"]
	fmt.Println(s)

	x := []int{3, 7, 2, 4, 3, 2, 3}
	goniq.RemoveDuplicates(&x) // returns [3, 7, 2, 4]
	goniq.Add(&x, 7) // returns [3, 7, 2, 4]
	goniq.Add(&x, 5) // returns [3, 7, 2, 4, 5]
	goniq.Remove(&x, 3) // returns [2, 4, 5, 7]
	fmt.Println(x)
}
```
