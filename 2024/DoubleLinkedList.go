package main

import "fmt"

// Node representa un nodo de la lista doblemente enlazada
type Node[T any] struct {
	Value T
	Prev  *Node[T]
	Next  *Node[T]
}

// DoublyLinkedList representa una lista doblemente enlazada
type DoublyLinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

// NewDoublyLinkedList crea una nueva lista doblemente enlazada
func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

// PushFront agrega un elemento al frente de la lista
func (dll *DoublyLinkedList[T]) PushFront(value T) {
	newNode := &Node[T]{Value: value}
	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
		dll.Size++
		return
	}
	newNode.Next = dll.Head
	dll.Head.Prev = newNode
	dll.Head = newNode
	dll.Size++
}

// PushBack agrega un elemento al final de la lista
func (dll *DoublyLinkedList[T]) PushBack(value T) {
	newNode := &Node[T]{Value: value}
	if dll.Tail == nil {
		dll.Head = newNode
		dll.Tail = newNode
		dll.Size++
		return
	}
	dll.Tail.Next = newNode
	newNode.Prev = dll.Tail
	dll.Tail = newNode
	dll.Size++
}

// PopFront elimina y retorna el elemento del frente de la lista
func (dll *DoublyLinkedList[T]) PopFront() (T, bool) {
	if dll.Head == nil {
		var zero T
		return zero, false
	}
	value := dll.Head.Value
	if dll.Head == dll.Tail {
		dll.Head = nil
		dll.Tail = nil
	} else {
		dll.Head = dll.Head.Next
		dll.Head.Prev = nil
	}
	dll.Size--
	return value, true
}

// PopBack elimina y retorna el elemento del final de la lista
func (dll *DoublyLinkedList[T]) PopBack() (T, bool) {
	if dll.Tail == nil {
		var zero T
		return zero, false
	}
	value := dll.Tail.Value
	if dll.Head == dll.Tail {
		dll.Head = nil
		dll.Tail = nil
	} else {
		dll.Tail = dll.Tail.Prev
		dll.Tail.Next = nil
	}
	dll.Size--
	return value, true
}

// InsertAfter inserta un nuevo nodo con un valor dado después de un nodo existente
func (dll *DoublyLinkedList[T]) InsertAfter(node *Node[T], value T) {
	if node == nil {
		return
	}
	newNode := &Node[T]{Value: value}
	newNode.Prev = node
	newNode.Next = node.Next
	if node.Next != nil {
		node.Next.Prev = newNode
	} else {
		dll.Tail = newNode
	}
	node.Next = newNode
	dll.Size++
}

// InsertBefore inserta un nuevo nodo con un valor dado antes de un nodo existente
func (dll *DoublyLinkedList[T]) InsertBefore(node *Node[T], value T) {
	if node == nil {
		return
	}
	newNode := &Node[T]{Value: value}
	newNode.Next = node
	newNode.Prev = node.Prev
	if node.Prev != nil {
		node.Prev.Next = newNode
	} else {
		dll.Head = newNode
	}
	node.Prev = newNode
	dll.Size++
}

// GetNext retorna el siguiente nodo de un nodo dado
func (dll *DoublyLinkedList[T]) GetNext(node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	}
	return node.Next
}

// GetPrev retorna el nodo anterior de un nodo dado
func (dll *DoublyLinkedList[T]) GetPrev(node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	}
	return node.Prev
}

func (dll *DoublyLinkedList[T]) GetHead() *Node[T] {
	return dll.Head
}

func (dll *DoublyLinkedList[T]) GetTail() *Node[T] {
	return dll.Tail
}

func (dll *DoublyLinkedList[T]) GetSize() int {
	return dll.Size
}

// PrintForward imprime los elementos de la lista en orden desde el frente
func (dll *DoublyLinkedList[T]) PrintForward() {
	for current := dll.Head; current != nil; current = current.Next {
		fmt.Printf("%v ", current.Value)
	}
	fmt.Println()
}

// PrintBackward imprime los elementos de la lista en orden inverso desde el final
func (dll *DoublyLinkedList[T]) PrintBackward() {
	for current := dll.Tail; current != nil; current = current.Prev {
		fmt.Printf("%v ", current.Value)
	}
	fmt.Println()
}

func main() {
	dll := NewDoublyLinkedList[int]()
	dll.PushBack(10)
	dll.PushBack(20)
	dll.PushFront(5)
	dll.PrintForward()  // Output: 5 10 20
	dll.PrintBackward() // Output: 20 10 5

	node := dll.Head.Next // Nodo con valor 10
	dll.InsertAfter(node, 15)
	dll.InsertBefore(node, 8)
	dll.PrintForward()  // Output: 5 8 10 15 20
	dll.PrintBackward() // Output: 20 15 10 8 5
}
