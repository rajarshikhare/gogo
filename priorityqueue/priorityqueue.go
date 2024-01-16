package priorityqueue

import "fmt"

type PriorityQueue[T any] struct {
	data       *[]T
	comparator func(T, T) bool
}

func New[T any](comparator func(T, T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{comparator: comparator, data: &[]T{}}
}

func (this *PriorityQueue[T]) Enqueue(val T) {
	*this.data = append(*this.data, val)
	currentIndex := len(*this.data) - 1
	for currentIndex > 0 {
		parent := currentIndex / 2
		if !this.comparator((*this.data)[parent], (*this.data)[currentIndex]) {
			(*this.data)[parent], (*this.data)[currentIndex] = (*this.data)[currentIndex], (*this.data)[parent]
			currentIndex = parent
		} else {
			break
		}
	}
}

func (this *PriorityQueue[T]) Dequeue() (T, bool) {
	if len(*this.data) > 0 {
		first := (*this.data)[0]
		(*this.data)[0] = (*this.data)[len(*this.data)-1]
		*this.data = (*this.data)[:len(*this.data)-1]
		for current := 0; current < len(*this.data); {
			leftIndex, rightIndex := current*2, current*2+1
			greaterNode := current

			if leftIndex < len(*this.data) {
				if this.comparator((*this.data)[leftIndex], (*this.data)[greaterNode]) {
					greaterNode = leftIndex
				}
			}

			if rightIndex < len(*this.data) {
				if this.comparator((*this.data)[rightIndex], (*this.data)[greaterNode]) {
					greaterNode = rightIndex
				}
			}

			if greaterNode != current {
				(*this.data)[current], (*this.data)[greaterNode] = (*this.data)[greaterNode], (*this.data)[current]
				current = greaterNode
			} else {
				break
			}

		}
		return first, true

	}
	var zeroValue T
	return zeroValue, false
}

func (this *PriorityQueue[T]) Peek(val T) (T, bool) {
	if len(*this.data) > 0 {
		first := (*this.data)[0]
		return first, true

	}
	var zeroValue T
	return zeroValue, false
}

func (this *PriorityQueue[T]) IsEmpty() bool {
	return !(len(*this.data) > 0)
}

func (this *PriorityQueue[T]) Size() int {
	return len(*this.data)
}

func (this *PriorityQueue[T]) Clear() {
	*this.data = []T{}
}

func (this *PriorityQueue[T]) String() string {
	res := "["
	for i := 0; i < len(*this.data); i++ {
		if i != len(*this.data)-1 {
			res = res + fmt.Sprintf("%v, ", (*this.data)[i])
		} else {
			res = res + fmt.Sprintf("%v", (*this.data)[i])
		}
	}
	res = res + "]"
	return res
}
