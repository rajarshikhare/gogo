## Overview

This package implements a generic queue in Go. A queue is a linear data structure that follows the First In, First Out (FIFO) principle. The package provides a straightforward and efficient way to use queues with any data type.

## Installation

To use the `queue` package, include it in your Go project:

```go
import "github.com/rajarshikhare/gogo/queue"
```

## Usage

The `queue` package offers the following functionalities:

### Type Declaration

```go
type Queue[T any] []T
```

### Constructor

- `New[T any]() *Queue[T]`: Creates and returns a new queue.

### Methods

- `Enqueue(val T)`: Adds an element to the end of the queue.
- `Dequeue() (T, bool)`: Removes and returns the first element of the queue. Returns `false` if the queue is empty.
- `Peek() (T, bool)`: Returns the first element without removing it. Returns `false` if the queue is empty.
- `IsEmpty() bool`: Checks if the queue is empty.
- `Size() int`: Returns the size of the queue.
- `Clear()`: Clears the queue.
- `String() string`: Returns a string representation of the queue.

## Example

Here is a simple example of how to use the `queue` package:

```go
package main

import (
	"fmt"
	"github.com/rajarshikhare/gogo/queue"
)

func main() {
	q := queue.New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q.String()) // Output: [1, 2]
	first, _ := q.Dequeue()
	fmt.Println(first)      // Output: 1
	fmt.Println(q.String()) // Output: [2]
}
```

## License

Specify your licensing information here.

## Contributions

Contributions are welcome. Please submit a pull request or open an issue for further discussion.

## Contact

For any queries or feedback, please contact the package maintainer.

---

**Note:** Replace `github.com/rajarshikhare/gogo/queue` with the actual path to the package in your project if it's different.