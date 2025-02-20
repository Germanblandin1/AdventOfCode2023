package main

import (
	"container/heap"
	"fmt"
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
func (pq *PriorityQueue[T]) Extract() (T, int, bool) {
	if pq.Len() == 0 {
		var zero T
		return zero, -1, false
	}
	item := heap.Pop(pq).(*Item[T])
	return item.Value, item.Priority, true
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

var INF = 100000000000

var n int
var m int
var mapa [][]rune
var dist [][]int
var iniI, iniJ, finI, finJ int
var movi = []int{0, -1, 0, 1}
var movj = []int{1, 0, -1, 0}

func PrintMapa() {
	for i := 0; i < n; i++ {
		fmt.Println(string(mapa[i]))
	}
}

type State struct {
	i, j int
}

func bfs(i, j, dir int) int {

	cola := make(PriorityQueue[State], 0)
	cola.Insert(State{i, j}, 0)
	dist[i][j] = 0
	for cola.Len() > 0 {
		actual, _, _ := cola.Extract()
		i := actual.i
		j := actual.j
		//fmt.Printf("i=%d j=%d dir=%d prio=%d dist=%d\n", i, j, dir, priority, dist[dir][i][j])
		if i == finI && j == finJ {
			return dist[i][j]
		}
		for k := 0; k < 4; k++ {
			ni := i + movi[k]
			nj := j + movj[k]
			if ni >= 0 && ni < n && nj >= 0 && nj < m && mapa[ni][nj] != '#' {
				if dist[ni][nj] > dist[i][j]+1 {
					dist[ni][nj] = dist[i][j] + 1
					cola.Insert(State{ni, nj}, dist[ni][nj])
				}
			}
		}
	}
	return -1

}

var points []*State

func f(lim int) bool {
	dist = make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			mapa[i][j] = '.'
			dist[i][j] = INF
		}
	}

	for i := 0; i < lim; i++ {
		mapa[points[i].j][points[i].i] = '#'
	}

	return bfs(iniI, iniJ, 0) > 0

}

func main() {

	//var total uint64 = 0
	tam := 25
	n = 7
	lim := 12
	tam = 3450
	n = 71
	lim = 1024
	mapa = make([][]rune, n)
	m = n
	for i := 0; i < n; i++ {
		mapa[i] = make([]rune, m)
		for j := 0; j < m; j++ {
			mapa[i][j] = '.'
		}
	}
	points = make([]*State, tam)
	for c := 0; c < tam; c++ {
		var i, j int
		fmt.Scanf("%d,%d\n", &i, &j)
		//fmt.Println("i=", i, "j=", j)
		estado := &State{i: i, j: j}
		//fmt.Println("estado=", estado)
		points[c] = estado
		//fmt.Println("points[", c, "]=", points[c])
	}

	for i := 0; i < lim; i++ {
		//fmt.Println("points[", i, "]=", points[i])
		mapa[points[i].j][points[i].i] = '#'
	}

	iniI, iniJ = 0, 0
	finI, finJ = n-1, n-1
	//PrintMapa()
	//busqueda binaria

	lo := 0
	hi := tam - 1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if !f(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	lo--
	fmt.Println(lo, hi)
	fmt.Printf("%d,%d\n", points[lo].i, points[lo].j)
}
