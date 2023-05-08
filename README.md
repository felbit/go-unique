# goniq

![Go Version](https://img.shields.io/badge/Go-v1.20-blue)
![Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)

*go unique* ... simple set-like operations for Go slices.

> :warning: This repository is archived and sunsetting because I started work on [set](github.com/felbit/go-set). `set` is basicall goniq2, with a different name. `goniq` stays here because some private projects depend on it.

## Goals

- "out of the way"
- generic
- idempotent

## Project Structure

`goniq` has two parts:

- Functions that operate on slices of `Ordered` types and provide handy set-like operations, like `RemoveDuplicates` to
  create a slice of unique elements, and `Add` and `Remove` that keep elements unique and work idempotent.
- The `Set` type is a thread-safe wrapper around these functions that provides a more convenient interface.

### Functions

#### Contains

```go
func Contains[T Ordered](s []T, e T) bool
```

This function takes a slice of T and an element of T as input, and returns true if the element is in the slice.

#### RemoveDuplicates

```go
func RemoveDuplicates[C comparable](*[]C)
```

This function takes a slice of a `comparable` type and returns possibly reduced set of unique entries of the original
slice.
Set is unsorted. The order of elements will be the first unique appearance of elements in the original slice.

#### Add

```go
func Add[T Ordered](*[]T, T)
```

This function takes a slice of T and an element of T as input,
and returns a new sorted slice of T with the input element added to it.
If the input slice is empty, it creates a new slice with only the input string.
If the input string is already in the slice, it returns the original slice, but sorted.

#### Remove

```go
func Remove[T Ordered](*[]T, T)
```

> :warning: `Remove` currently returns a **sorted slice**, even if the elements fo not change. This might change in the
> future.

This function takes a slice of T and an element of T as input,
and returns a new sorted slice of T with the input element removed from it.
If the input slice is empty or the input element is not in the slice,
it returns the original slice sorted.

### Set

`Set` is a thread-safe wrapper around the functions above. It provides a more convenient interface and some additional
methods.

## Usage

```go
package main

import (
	"fmt"
	"github.com/felbit/goniq"
)

func main() {
	s := []string{"apple", "banana", "apple", "cherry"}
	goniq.RemoveDuplicates(&s) // ["apple", "banana", "cherry"]
	goniq.Add(&s, "banana")    // ["apple", "banana", "cherry"]
	goniq.Add(&s, "date")      // ["apple", "banana", "cherry", "date"]
	goniq.Remove(&s, "banana") // ["apple", "cherry", "date"]

	set := goniq.NewSet[int]()
	set.Add(3)                      // [3]
	set.Append(3, 7, 2, 4, 3, 2, 3) // [3, 7, 2, 4]
	set.Contains(3)                 // true
	set.Add(7)                      // [3, 7, 2, 4]
	set.Remove(3)                   // [2, 4, 7]
	set.RemoveAll()                 // []
}
```
