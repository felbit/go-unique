# goniq
![Go Version](https://img.shields.io/badge/Go-v1.20-blue)
![Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)

*go unique* ... simple set-like operations for Go slices.

>   :warning: This library is not considered stable. Therefore, caution is advised using it for production for now.

Goniq is a Go language package that provides functions to add and remove elements from a (sorted) string slice.

## Goals

- "out of the way": using this package should be intuitively and not require reading documentation
- generic implementation: where-ever possible, users should not have to care about the types of things
- idempotent behavior: running the same operation multiple times should not change the result (this is currently violated by the behavior of `Remove` on slices with non-unique entries)
- independent: no use of external libraries
- resource conscious: use as little resources as possible

## Functions

### Add

```go
func Add[T Ordered]([]T, T) []T
```

>   :warning: `Add` currently returns a **sorted slice**, even if the elements do not change. This might change in the future.

This function takes a slice of T and an element of T as input,
and returns a new sorted slice of T with the input element added to it.
If the input slice is empty, it creates a new slice with only the input string.
If the input string is already in the slice, it returns the original slice, but sorted.


### Remove

```go
func Remove[T Ordered]([]T, T) []T
```

>   :warning: `Remove` currently returns a **sorted slice**, even if the elements fo not change. This might change in the future.

>   :warning: Currently, only removes one of the entries, if there are multiple entries that match the input element! This will change in the future.

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
	s := []string{"apple", "banana", "cherry"}
	s = goniq.Add(s, "banana") // returns ["apple", "banana", "cherry"]
	s = goniq.Add(s, "date") // returns ["apple", "banana", "cherry", "date"]
	s = goniq.Remove(s, "banana") // returns ["apple", "cherry", "date"]
	fmt.Println(s)

	x := []int{3, 7, 2, 4}
	x = goniq.Add(x, 7) // returns [2, 3, 4, 7]
	x = goniq.Add(x, 5) // returns [2, 3, 4, 5, 7]
	x = goniq.Remove(x, 3) // returns [2, 4, 5, 7]
	fmt.Println(x)
}
```
