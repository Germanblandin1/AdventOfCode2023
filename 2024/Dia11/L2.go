package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

var graph map[string][]string
var visited map[string]bool
var n int
var maximaaltura int
var stonesOri []string

func extraerHijos(value string) []string {
	resp := make([]string, 0)
	if value == "0" {
		resp = append(resp, "1")
	} else if len(value)%2 == 0 {
		//partimos a la mitad el string y lo agregamos a la lista
		mitad := len(value) / 2
		left := value[:mitad]
		right := value[mitad:]
		uintLeft, _ := strconv.ParseUint(left, 10, 64)
		uintRight, _ := strconv.ParseUint(right, 10, 64)

		left = strconv.FormatUint(uintLeft, 10)
		right = strconv.FormatUint(uintRight, 10)

		resp = append(resp, left)
		resp = append(resp, right)
	} else {
		// convertimos el valor a entero y lo multiplicamos por 2024 y sustituimos el valor
		uintValue, _ := strconv.ParseUint(value, 10, 64)
		newValue := strconv.FormatUint(uintValue*2024, 10)
		resp = append(resp, newValue)
	}
	return resp
}

type Key struct {
	node   string
	altura int
}

var memo map[Key]uint64
var mark map[Key]bool

func dfsCount(key Key) uint64 {

	if key.altura == maximaaltura {
		return 1
	}

	if mark[key] {
		return memo[key]
	}

	mark[key] = true

	memo[key] = 0
	for _, child := range graph[key.node] {
		memo[key] += dfsCount(Key{child, key.altura + 1})
	}
	return memo[key]
}

func main() {

	var total uint64 = 0
	tam := 1
	//tam = 53
	veces := 75
	reader := bufio.NewReader(os.Stdin)
	stonesOri = make([]string, 0)
	for c := 0; c < tam; c++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		strStones := strings.Split(line, " ")
		n = len(strStones)
		graph = make(map[string][]string)
		visited = make(map[string]bool)
		cola := NewDoublyLinkedList[string]()
		for i := 0; i < n; i++ {
			stonesOri = append(stonesOri, strStones[i])
			cola.PushBack(strStones[i])
		}
		for i := 0; i < veces; i++ {
			for cola.GetSize() > 0 {
				head, _ := cola.PopFront()
				//fmt.Println("head", head)
				if !visited[head] {
					hijos := extraerHijos(head)
					graph[head] = make([]string, 0)
					graph[head] = append(graph[head], hijos...)
					visited[head] = true
					for _, hijo := range hijos {
						if !visited[hijo] {
							cola.PushBack(hijo)
						}
					}
				}
			}
		}
	}
	maximaaltura = veces

	memo = make(map[Key]uint64)
	mark = make(map[Key]bool)
	for _, stone := range stonesOri {
		total += dfsCount(Key{stone, 0})
	}
	fmt.Println(total)
}
