package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type Node[T any] struct {
	pre *Node[T]
	nxt *Node[T]
	val T
}
type MyList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func CreateList[T any](arr []T) *MyList[T] {
	Li := &MyList[T]{}
	for _, v := range arr {
		Li = Li.Add(v)
	}
	return Li
}

func (m *MyList[T]) Add(v T) *MyList[T] {

	np := &Node[T]{nil, nil, v}
	if m.head == nil {
		m.head = np
		m.tail = np
	}

	np.pre = m.tail
	m.tail.nxt = np
	m.tail = np
	return m
}
func (m *MyList[T]) Remov() *MyList[T] {
	m.tail = m.tail.pre
	m.tail.nxt = nil
	return m
}
func (m *MyList[T]) Display() {
	np := m.head
	for np != nil {
		fmt.Printf("%v", np.val)
		np = np.nxt
	}

	fmt.Println()
}
func main() {
	arr := []int{1, 2, 3, 4}
	List := CreateList(arr)
	List.Display()
	List.Add(9)
	List.Display()
	List.Remov()
	List.Display()
}
