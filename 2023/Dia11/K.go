package main

import (
	"container/heap"
	"fmt"
)

var basura string
var t int
var n, m int
var INF int = 1000000000

var matrizORI [][]rune
var matriz [][]rune

var movi = []int{-1, 0, 1, 0}
var movj = []int{0, 1, 0, -1}

// arriba:0 derecha:1 abajo:2 izquierda:3

var distancias [][]int
var marca [][]int

type Pair struct {
	i, j int
}

type Queue []Pair

func (q *Queue) Enqueue(p Pair) {
	*q = append(*q, p)
}

func (q *Queue) Dequeue() (Pair, bool) {
	if len(*q) == 0 {
		return Pair{}, false
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, true
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

type Item struct {
	value    Pair
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) IsEmpty() bool {
	return len(*pq) == 0
}

func bfs(i, j int) int {

	distancias = make([][]int, n)
	for i := 0; i < n; i++ {
		distancias[i] = make([]int, m)
		for j := 0; j < m; j++ {
			distancias[i][j] = INF
		}
	}

	distancias[i][j] = 0
	cola := PriorityQueue{}
	heap.Init(&cola)
	item := &Item{
		value:    Pair{i, j},
		priority: 0,
	}
	heap.Push(&cola, item)
	suma := 0
	for !cola.IsEmpty() {
		it := heap.Pop(&cola).(*Item)
		pair := it.value
		//prioridad := it.priority
		i := pair.i
		j := pair.j
		//fmt.Println(i, j, prioridad)
		if matriz[i][j] == '#' {
			suma += distancias[i][j]
		}
		for k := 0; k < 4; k++ {
			ni := i + movi[k]
			nj := j + movj[k]
			if ni >= 0 && ni < n && nj >= 0 && nj < m {
				value := 1
				if repiteCol[nj] || repiteFila[ni] {
					value = 1000000
				}

				if distancias[i][j]+value < distancias[ni][nj] {
					distancias[ni][nj] = distancias[i][j] + value
					item := &Item{
						value:    Pair{ni, nj},
						priority: distancias[ni][nj],
					}
					heap.Push(&cola, item)
				}
			}
		}
	}
	return suma

}

var repiteFila []bool
var repiteCol []bool

func main() {

	fmt.Scan(&n)
	fmt.Scan(&m)
	repiteFila = make([]bool, n)
	repiteCol = make([]bool, m)
	for i := 0; i < n; i++ {
		repiteFila[i] = true
	}
	for j := 0; j < m; j++ {
		repiteCol[j] = true
	}

	matriz = make([][]rune, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&basura)
		matriz[i] = []rune(basura)
		for j := 0; j < m; j++ {
			if matriz[i][j] == '#' {
				repiteFila[i] = false
				repiteCol[j] = false
			}
		}
	}

	canF := 0
	for i := 0; i < n; i++ {
		if repiteFila[i] {
			canF++
		}
	}
	canC := 0
	for j := 0; j < m; j++ {
		if repiteCol[j] {
			canC++
		}
	}

	fmt.Println(n, m)
	for i := 0; i < n; i++ {
		fmt.Println(string(matriz[i]))
	}

	total := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matriz[i][j] == '#' {
				total += bfs(i, j)
			}
		}
	}

	fmt.Println(total / 2)

}
