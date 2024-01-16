## Overview

The GoGo Data Structures Package is a comprehensive collection of generic data structures implemented in Go. This package includes implementations of Stack, Queue, and PriorityQueue, each designed to be efficient, flexible, and easy to use. These data structures are fundamental in many algorithms and are essential for efficient data management and processing.

## Data Structures

### 1. Stack
   - **Package Path**: `github.com/rajarshikhare/gogo/stack`
   - **Description**: A Last In, First Out (LIFO) data structure.
   - **Key Methods**: `Push`, `Pop`, `Top`, `IsEmpty`, `Size`, `Clear`, `String`.
   - **Usage**: Ideal for scenarios where you need to access the most recently added elements first.

### 2. Queue
   - **Package Path**: `github.com/rajarshikhare/gogo/queue`
   - **Description**: A First In, First Out (FIFO) data structure.
   - **Key Methods**: `Enqueue`, `Dequeue`, `Peek`, `IsEmpty`, `Size`, `Clear`, `String`.
   - **Usage**: Suitable for situations where you need to process elements in the order they were added.

### 3. PriorityQueue
   - **Package Path**: `github.com/rajarshikhare/gogo/priorityqueue`
   - **Description**: An advanced queue where each element is associated with a priority.
   - **Key Methods**: `Enqueue`, `Dequeue`, `Peek`, `IsEmpty`, `Size`, `Clear`, `String`.
   - **Usage**: Perfect for scenarios where elements need to be processed based on priority rather than just the order of addition.

## Installation

To use any of these data structures in your Go project, import the respective package:

```go
import "github.com/rajarshikhare/gogo/[data_structure]"
```

Replace `[data_structure]` with `stack`, `queue`, or `priorityqueue` as needed.

## Examples

### Stack
```go
s := stack.New[int]()
s.Push(10)
fmt.Println(s.Pop())
```

### Queue
```go
q := queue.New[int]()
q.Enqueue(10)
fmt.Println(q.Dequeue())
```

### PriorityQueue
```go
pq := priorityqueue.New[int](func(a, b int) bool { return a < b })
pq.Enqueue(10)
fmt.Println(pq.Dequeue())
```

## Contributing

Contributions to enhance the functionalities, performance, or documentation of these data structures are welcome. Please feel free to submit a pull request or open an issue for discussion.

## License

Specify your licensing information here.

## Contact

For any queries, suggestions, or feedback regarding this package, please reach out to the package maintainer.

---