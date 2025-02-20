package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
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
var dist [][][]int
var visited [][][]bool
var iniI, iniJ int
var finI, finJ int

var movi = []int{0, -1, 0, 1}
var movj = []int{1, 0, -1, 0}
var movc = map[rune]int{'>': 0, '^': 1, '<': 2, 'v': 3}

func PrintMapa() {
	for i := 0; i < n; i++ {
		fmt.Println(string(mapa[i]))
	}
}

type State struct {
	i, j, dir int
}

func bfs(i, j, dir int) int {
	dist = make([][][]int, 4)
	for i := 0; i < 4; i++ {
		dist[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = make([]int, m)
			for k := 0; k < m; k++ {
				dist[i][j][k] = INF
			}
		}
	}

	cola := make(PriorityQueue[State], 0)
	cola.Insert(State{i, j, dir}, 0)
	dist[dir][i][j] = 0
	for cola.Len() > 0 {
		actual, _, _ := cola.Extract()
		i := actual.i
		j := actual.j
		dir := actual.dir
		//fmt.Printf("i=%d j=%d dir=%d prio=%d dist=%d\n", i, j, dir, priority, dist[dir][i][j])
		if i == finI && j == finJ {
			return dist[dir][i][j]
		}
		for k := -1; k <= 1; k++ {
			newDir := (dir + k + 4) % 4
			ni := i
			nj := j
			sum := 1000
			if dir == newDir {
				ni = i + movi[newDir]
				nj = j + movj[newDir]
				sum = 1
			}

			//fmt.Printf("Posible i=%d j=%d dir=%d\n", ni, nj, newDir)
			if ni < 0 || ni >= n || nj < 0 || nj >= m {
				continue
			}
			if mapa[ni][nj] == '#' {
				continue
			}
			if dist[dir][i][j]+sum < dist[newDir][ni][nj] {
				dist[newDir][ni][nj] = dist[dir][i][j] + sum
				//fmt.Printf("encolo i=%d j=%d dir=%d prio=%d dist=%d\n", ni, nj, newDir, dist[newDir][ni][nj], dist[dir][i][j]+sum)
				cola.Insert(State{ni, nj, newDir}, dist[newDir][ni][nj])
			}
		}
	}
	return -1

}

func main() {

	var total uint64 = 0
	tam := 17
	tam = 141
	n = tam
	mapa = make([][]rune, n)

	reader := bufio.NewReader(os.Stdin)
	for c := 0; c < tam; c++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		mapa[c] = []rune(line)
		m = len(mapa[c])
		for j := 0; j < m; j++ {
			if mapa[c][j] == 'S' {
				iniI = c
				iniJ = j
				mapa[c][j] = '.'
			}
			if mapa[c][j] == 'E' {
				finI = c
				finJ = j
				mapa[c][j] = '.'
			}
		}
	}

	total = uint64(bfs(iniI, iniJ, 0))

	fmt.Println(total)
}
