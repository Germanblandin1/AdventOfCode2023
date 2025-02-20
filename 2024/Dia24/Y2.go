package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

// Item representa un elemento en la cola de prioridad.
type Item[T any] struct {
	Value    T   // El valor almacenado en el elemento.
	Priority int // La prioridad del elemento. Un número menor indica mayor prioridad.
}

// PriorityQueue es una cola de prioridad genérica basada en un heap.
type PriorityQueue[T any] []*Item[T]

// Len devuelve la cantidad de elementos en la cola.
func (pq PriorityQueue[T]) Len() int { return len(pq) }

// Less compara las prioridades de dos elementos.
// Nota: Cambia la comparación si prefieres una cola de prioridad máxima.
func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority // Menor prioridad tiene más prioridad.
}

// Swap intercambia dos elementos en la cola.
func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push agrega un elemento a la cola. Este método es utilizado por `container/heap`.
func (pq *PriorityQueue[T]) Push(x any) {
	item := x.(*Item[T])
	*pq = append(*pq, item)
}

// Pop elimina y devuelve el elemento con mayor prioridad. Este método es utilizado por `container/heap`.
func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Insert agrega un elemento a la cola de prioridad.
func (pq *PriorityQueue[T]) Insert(value T, priority int) {
	heap.Push(pq, &Item[T]{Value: value, Priority: priority})
}

// Extract devuelve el elemento con mayor prioridad y lo elimina de la cola.
// Si la cola está vacía, devuelve el valor por defecto de `T` y `false`.
func (pq *PriorityQueue[T]) Extract() (T, bool) {
	if pq.Len() == 0 {
		var zero T
		return zero, false
	}
	item := heap.Pop(pq).(*Item[T])
	return item.Value, true
}

// Peek devuelve el elemento con mayor prioridad sin eliminarlo.
// Si la cola está vacía, devuelve el valor por defecto de `T` y `false`.
func (pq *PriorityQueue[T]) Peek() (T, bool) {
	if pq.Len() == 0 {
		var zero T
		return zero, false
	}
	return (*pq)[0].Value, true
}

// String devuelve un string con todos los elementos de la cola de prioridad.
func (pq *PriorityQueue[T]) String() string {
	result := "PriorityQueue:\n"
	for _, item := range *pq {
		result += fmt.Sprintf("Value: %v, Priority: %d\n", item.Value, item.Priority)
	}
	return result
}

var n int
var subn int
var mapeo map[string]int
var mapeo2 map[int]string
var grafo [][]int

// op 0 = 0
// op 1 = 1
// op 2 = and
// op 3 = or
// op 4 = xor
var estado []int
var estadoOriginal []int

func armar(node string) string {
	if mapeo[node] < subn {
		return node
	}
	ai := mapeo[node]
	a := armar(mapeo2[grafo[ai][0]])
	b := armar(mapeo2[grafo[ai][1]])
	if b < a {
		a, b = b, a
	}
	op := estado[ai]
	if op == 0 {
		return "0"
	}
	if op == 1 {
		return "1"
	}
	if op == 2 {
		return "(" + a + " AND " + b + ")"
	}
	if op == 3 {
		return "(" + a + " OR " + b + ")"
	}
	if op == 4 {
		return "(" + a + " XOR " + b + ")"
	}
	return ""
}

func addNodo(a string) int {
	ai, ok := mapeo[a]
	if !ok {
		mapeo[a] = n
		mapeo2[n] = a
		ai = n
		grafo = append(grafo, make([]int, 0))
		estado = append(estado, 0)
		n++
	}
	return ai
}

func main() {
	var total uint64 = 0
	tam := 10
	tam = 12
	tam = 90
	tam2 := 36
	tam2 = 6
	tam2 = 222
	n = 0

	subn = tam
	//reader := bufio.NewReader(os.Stdin)
	grafo = make([][]int, 0)
	mapeo = make(map[string]int)
	mapeo2 = make(map[int]string)
	estado = make([]int, 0)

	for c := 0; c < tam; c++ {
		var a string
		var num int
		fmt.Scanf("%s %d\n", &a, &num)
		a = strings.Trim(a, " ")
		a = strings.Trim(a, ":")

		ai := addNodo(a)
		estado[ai] = num
	}

	for c := 0; c < tam2; c++ {
		var a, b, op, c string
		fmt.Scanf("%s %s %s -> %s\n", &a, &op, &b, &c)
		a1 := addNodo(a)
		b1 := addNodo(b)
		c1 := addNodo(c)

		opInt := 0
		if op == "AND" {
			opInt = 2
		}
		if op == "OR" {
			opInt = 3
		}
		if op == "XOR" {
			opInt = 4
		}
		grafo[c1] = append(grafo[c1], a1)
		grafo[c1] = append(grafo[c1], b1)
		estado[c1] = opInt
	}

	nodos := make([]string, 0)
	for c := 0; c < n; c++ {
		if mapeo2[c][0] == 'z' {
			nodos = append(nodos, mapeo2[c])
		}
	}
	sort.Strings(nodos)

	for _, nodo := range nodos {
		ai := mapeo[nodo]
		fmt.Println("ai: ", nodo)
		expression := armar(mapeo2[ai])
		fmt.Println("expression: ", expression)

	}

	fmt.Println(total)

}
