package queue

import "fmt"

type Queue[T any] []T

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (this *Queue[T]) Enqueue(val T) {
	*this = append(*this, val)
}

func (this *Queue[T]) Dequeue() (T, bool) {
	if len(*this) > 0 {
		first := (*this)[0]
		*this = (*this)[1:]
		return first, true

	}
	var zeroValue T
	return zeroValue, false
}

func (this *Queue[T]) Peek(val T) (T, bool) {
	if len(*this) > 0 {
		first := (*this)[0]
		return first, true

	}
	var zeroValue T
	return zeroValue, false
}

func (this *Queue[T]) IsEmpty() bool {
	return !(len(*this) > 0)
}

func (this *Queue[T]) Size() int {
	return len(*this)
}

func (this *Queue[T]) Clear() {
	*this = []T{}
}

func (this *Queue[T]) String() string {
	res := "["
	for i := 0; i < len(*this); i++ {
		if i != len(*this)-1 {
			res = res + fmt.Sprintf("%v, ", (*this)[i])
		} else {
			res = res + fmt.Sprintf("%v", (*this)[i])
		}
	}
	res = res + "]"
	return res
}
