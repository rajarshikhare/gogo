package main

import (
	"fmt"

	"github.com/rajarshikhare/gogo/priorityqueue"
	"github.com/rajarshikhare/gogo/queue"
	"github.com/rajarshikhare/gogo/stack"
)

func main() {
	a := queue.New[string]()
	a.Enqueue("1")
	fmt.Println(a)
	a.Dequeue()
	fmt.Println(a)
	a.Enqueue("2")
	fmt.Println(a)
	a.Enqueue("3")
	fmt.Println(a)
	a.Dequeue()
	fmt.Println(a)
	a.Enqueue("4")
	fmt.Println(a)
	a.Enqueue("5")
	fmt.Println(a)

	s := stack.New[int]()
	s.Push(1)
	fmt.Println(s)
	s.Push(2)
	fmt.Println(s)
	s.Push(3)
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)

	p := priorityqueue.New[string](func(a, b string) bool {
		return a > b
	})

	p.Enqueue("sdf")
	fmt.Println(p)
	p.Enqueue("sdf")
	fmt.Println(p)
	p.Enqueue("3")
	fmt.Println(p)
	p.Enqueue("")
	fmt.Println(p)
	p.Enqueue("100")
	fmt.Println(p)
	p.Enqueue("50")
	fmt.Println(p)

	//
	zz, _ := p.Dequeue()
	fmt.Println(zz)
	zz, _ = p.Dequeue()
	fmt.Println(zz)
	zz, _ = p.Dequeue()
	fmt.Println(zz)
	zz, _ = p.Dequeue()
	fmt.Println(zz)
	zz, _ = p.Dequeue()
	fmt.Println(zz)
	zz, _ = p.Dequeue()
	fmt.Println(zz)
	zz, o := p.Dequeue()
	fmt.Println(o)

}
