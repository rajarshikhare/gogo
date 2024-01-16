package stack

import "fmt"

type Stack[T any] []T

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (this *Stack[T]) Push(val T) {
	*this = append(*this, val)
}

func (this *Stack[T]) Pop() (T, bool) {
	if len(*this) > 0 {
		top := (*this)[len(*this)-1]
		*this = (*this)[:len(*this)-1]
		return top, true

	}
	var zeroValue T
	return zeroValue, false
}

func (this *Stack[T]) Top(val T) (T, bool) {
	if len(*this) > 0 {
		first := (*this)[0]
		return first, true

	}
	var zeroValue T
	return zeroValue, false
}

func (this *Stack[T]) IsEmpty() bool {
	return !(len(*this) > 0)
}

func (this *Stack[T]) Size() int {
	return len(*this)
}

func (this *Stack[T]) Clear() {
	*this = []T{}
}

func (this *Stack[T]) String() string {
	res := "["
	for i := 0; i < len(*this); i++ {
		if i != len(*this)-1 {
			res = res + fmt.Sprintf("%v, ", (*this)[i])
		} else {
			res = res + fmt.Sprintf("%v", (*this)[i])
		}
	}
	res = res + "] <- Top"
	return res
}
