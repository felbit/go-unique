# goniq
*go unique* ... simple set-like operations for Go slices.

:warning: This library is not considered stable. Therefore caution is advised using it for production right now.

GoUnique is a Go language package that provides functions to add and remove elements from a (sorted) string slice. The package contains the functions `Add` and `Remove`.

## Functions

### Add

```go
func Add([]string, string) []string
```

:warning: `Add` currently returns a **sorted slice**. Event if the elements do not change, the slice will be modified!

This function takes a string slice and a string as input, and returns a new sorted string slice with the input string added to it. If the input slice is empty, it creates a new slice with only the input string. If the input string is already in the slice, it returns the original slice without modifying it. Otherwise, it uses the sort.SearchStrings function from the standard library to find the position where the input string should be inserted, and inserts it using the append function.


### Remove

```go
func Remove([]string, string) []string
```

:warning: `Remove` currently returns a **sorted slice**. Event if the elements do not change, the slice will be modified!

This function takes a string slice and a string as input, and returns a new sorted string slice with the input string removed from it. If the input slice is empty or the input string is not in the slice, it returns the original slice it. Otherwise, it uses the sort.SearchStrings function to find the position of the input string in the slice, and removes it using the copy function.

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
}
```
