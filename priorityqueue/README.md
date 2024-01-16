## Overview

The `priorityqueue` package in Go provides an implementation of a priority queue. A priority queue is a data structure where each element has a priority associated with it. Elements with higher priority are served before elements with lower priority. This package offers a customizable priority mechanism based on a user-defined comparator function.

## Installation

To use the `priorityqueue` package, include it in your Go project:

```go
import "github.com/rajarshikhare/gogo/priorityqueue"
```

## Usage

The `priorityqueue` package includes the following key components:

### Type Declaration

```go
type PriorityQueue[T any] struct {
	data       *[]T
	comparator func(T, T) bool
}
```

### Constructor

- `New[T any](comparator func(T, T) bool) *PriorityQueue[T]`: Creates a new priority queue with the specified comparator function.

### Methods

- `Enqueue(val T)`: Adds an element to the priority queue.
- `Dequeue() (T, bool)`: Removes and returns the element with the highest priority. Returns `false` if the queue is empty.
- `Peek() (T, bool)`: Returns the highest-priority element without removing it. Returns `false` if the queue is empty.
- `IsEmpty() bool`: Checks if the priority queue is empty.
- `Size() int`: Returns the size of the priority queue.
- `Clear()`: Clears the priority queue.
- `String() string`: Returns a string representation of the priority queue.

## Example

Here is an example of how to use the `priorityqueue` package:

```go
package main

import (
	"fmt"
	"github.com/rajarshikhare/gogo/priorityqueue"
)

func main() {
	// Comparator function for integer priority
	comparator := func(a, b int) bool {
		return a < b // Lower number, higher priority
	}

	pq := priorityqueue.New[int](comparator)
	pq.Enqueue(5)
	pq.Enqueue(1)
	pq.Enqueue(3)
	fmt.Println(pq.String()) // Output: [1, 5, 3]

	highest, _ := pq.Dequeue()
	fmt.Println(highest)     // Output: 1
	fmt.Println(pq.String()) // Output: [3, 5]
}
```

## License

Specify your licensing information here.

## Contributions

Contributions are welcome. Please submit a pull request or open an issue for further discussion.

## Contact

For any queries or feedback, please contact the package maintainer.

---

**Note:** Replace `github.com/rajarshikhare/gogo/priorityqueue` with the actual path to the package in your project if it's different.