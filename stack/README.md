Overview

This package implements a generic stack in Go. A stack is a linear data structure that follows the Last In, First Out (LIFO) principle. The package provides a simple and efficient way to use stacks with any data type.

Installation

To use the stack package, include it in your Go project:

go
Copy code
import "github.com/rajarshikhare/gogo/stack"
Usage

The stack package provides the following functionalities:

Type Declaration
go
Copy code
type Stack[T any] []T
Constructor
New[T any]() *Stack[T]: Creates and returns a new stack.
Methods
Push(val T): Adds an element to the top of the stack.
Pop() (T, bool): Removes and returns the top element of the stack. Returns false if the stack is empty.
Top(val T) (T, bool): Returns the top element without removing it. Returns false if the stack is empty.
IsEmpty() bool: Checks if the stack is empty.
Size() int: Returns the size of the stack.
Clear(): Clears the stack.
String() string: Returns a string representation of the stack, with the top element on the right.
Example

Here is a simple example of how to use the stack package:

go
Copy code
package main

import (
	"fmt"
	"github.com/rajarshikhare/gogo/stack"
)

func main() {
	s := stack.New[int]()
	s.Push(1)
	s.Push(2)
	fmt.Println(s.String()) // Output: [1, 2] <- Top
	top, _ := s.Pop()
	fmt.Println(top)        // Output: 2
	fmt.Println(s.String()) // Output: [1] <- Top
}
License

Specify your licensing information here.

Contributions

Contributions are welcome. Please submit a pull request or open an issue for further discussion.

Contact

For any queries or feedback, please contact the package maintainer.