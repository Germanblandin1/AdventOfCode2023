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

type Pair struct {
	i, j int
}

var movi = []int{0, 0, 1, -1}
var movj = []int{1, -1, 0, 0}

var dist [][]int
var mapa [][]rune
var n int
var m int
var iniI, iniJ int
var finI, finJ int
var INF = 1000000000

func dijkstra() int {
	fmt.Println("n", n, "m", m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dist[i][j] = INF
		}
	}

	pq := make(PriorityQueue[Pair], 0)
	pq.Insert(Pair{iniI, iniJ}, 0)
	dist[iniI][iniJ] = 0

	for pq.Len() > 0 {
		pair, _, _ := pq.Extract()
		i := pair.i
		j := pair.j

		if i == finI && j == finJ {
			return dist[i][j]
		}

		for k := 0; k < 4; k++ {
			ni := i + movi[k]
			nj := j + movj[k]
			if ni >= 0 && ni < n && nj >= 0 && nj < m && mapa[ni][nj] != '#' {
				if dist[ni][nj] > dist[i][j]+1 {
					dist[ni][nj] = dist[i][j] + 1
					pq.Insert(Pair{ni, nj}, dist[ni][nj])
				}
			}
		}

	}

	return INF
}

func main() {
	var total uint64 = 0
	tam := 15
	tam = 141
	dif := 100
	reader := bufio.NewReader(os.Stdin)

	n = tam
	mapa = make([][]rune, n)
	dist = make([][]int, n)
	for c := 0; c < tam; c++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		mapa[c] = []rune(line)
		m = len(mapa[c])
		dist[c] = make([]int, m)
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

	fmt.Println(iniI, iniJ, finI, finJ)
	totalValue := dijkstra()
	fmt.Println(totalValue)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mapa[i][j] == '#' {
				for k := 0; k < 4; k++ {
					for l := 0; l < 4; l++ {
						if k == l {
							continue
						}
						ni := i + movi[k]
						nj := j + movj[k]
						nni := i + movi[l]
						nnj := j + movj[l]
						if ni >= 0 && ni < n && nj >= 0 && nj < m {
							if nni >= 0 && nni < n && nnj >= 0 && nnj < m {
								if mapa[ni][nj] == '.' && mapa[nni][nnj] == '.' {
									newVal := dist[ni][nj] + 2 + (totalValue - dist[nni][nnj])
									if newVal < totalValue && totalValue-newVal >= dif {
										total++
									}
								}
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(total)
}
