package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
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

var INF int = 0x3f3f3f3f

var movi = []int{-1, 0, 1, 0}
var movj = []int{0, -1, 0, 1}

var tecladoNum [][]int = [][]int{
	{7, 8, 9},
	{4, 5, 6},
	{1, 2, 3},
	{-1, 0, 10},
}

var tecladoDir [][]int = [][]int{
	{-1, 0, 10},
	{1, 2, 3},
}

type Estado struct {
	ni, nj   int
	t1i, t1j int
	t2i, t2j int
	pos      int
}

var dist map[Estado]int
var finalCode string

func Equal(a, b Estado) bool {
	return a.ni == b.ni && a.nj == b.nj && a.t1i == b.t1i && a.t1j == b.t1j && a.t2i == b.t2i && a.t2j == b.t2j && a.pos == b.pos
}

func newEstado(estado Estado, pi, pj int) Estado {
	nEstado := Estado{
		ni:  estado.ni,
		nj:  estado.nj,
		t1i: estado.t1i,
		t1j: estado.t1j,
		t2i: estado.t2i,
		t2j: estado.t2j,
		pos: estado.pos,
	}

	if tecladoDir[pi][pj] == 10 {
		if tecladoDir[estado.t2i][estado.t2j] == 10 {
			if tecladoDir[estado.t1i][estado.t1j] == 10 {
				if tecladoNum[estado.ni][estado.nj] == 10 {
					if finalCode[estado.pos] == 'A' {
						nEstado.pos++
					}
				} else {
					if finalCode[estado.pos] == byte(tecladoNum[estado.ni][estado.nj]+'0') {
						nEstado.pos++
					}
				}

			} else {
				ni, nj := estado.ni+movi[tecladoDir[estado.t1i][estado.t1j]], estado.nj+movj[tecladoDir[estado.t1i][estado.t1j]]
				if ni >= 0 && ni < 4 && nj >= 0 && nj < 3 && tecladoNum[ni][nj] != -1 {
					nEstado.ni, nEstado.nj = ni, nj
				}
			}

		} else {
			ni, nj := estado.t1i+movi[tecladoDir[estado.t2i][estado.t2j]], estado.t1j+movj[tecladoDir[estado.t2i][estado.t2j]]
			if ni >= 0 && ni < 2 && nj >= 0 && nj < 3 && tecladoDir[ni][nj] != -1 {
				nEstado.t1i, nEstado.t1j = ni, nj
			}
		}

	} else {
		ni, nj := estado.t2i+movi[tecladoDir[pi][pj]], estado.t2j+movj[tecladoDir[pi][pj]]
		if ni >= 0 && ni < 2 && nj >= 0 && nj < 3 && tecladoDir[ni][nj] != -1 {
			nEstado.t2i, nEstado.t2j = ni, nj
		}
	}
	return nEstado
}

func disjktra(estadoIni Estado) int {

	estadoFin := Estado{
		ni:  3,
		nj:  2,
		t1i: 0,
		t1j: 2,
		t2i: 0,
		t2j: 2,
		pos: len(finalCode),
	}
	dist = make(map[Estado]int)
	//mark = make(map[Estado]bool)
	pq := make(PriorityQueue[Estado], 0)

	dist[estadoIni] = 0
	pq.Insert(estadoIni, 0)

	for pq.Len() > 0 {
		estado, _, _ := pq.Extract()
		if Equal(estado, estadoFin) {
			return dist[estado]
		}

		for i := 0; i < 2; i++ {
			for j := 0; j < 3; j++ {
				if i == 0 && j == 0 {
					continue
				}
				nEstado := newEstado(estado, i, j)

				dis, ok := dist[nEstado]
				if !ok || dis > dist[estado]+1 {
					dist[nEstado] = dist[estado] + 1
					pq.Insert(nEstado, dist[nEstado])
				}
			}
		}
	}
	return INF
}

func main() {
	var total uint64 = 0
	tam := 5
	reader := bufio.NewReader(os.Stdin)

	for c := 0; c < tam; c++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		finalCode = line
		estadoIni := Estado{
			ni:  3,
			nj:  2,
			t1i: 0,
			t1j: 2,
			t2i: 0,
			t2j: 2,
			pos: 0,
		}
		dist = make(map[Estado]int)
		fmt.Println(finalCode)
		val := disjktra(estadoIni)
		fmt.Println("--", val)

		finalCodeInt, _ := strconv.Atoi(finalCode[0 : len(finalCode)-1])
		fmt.Println("--", finalCodeInt)
		total += uint64(val) * uint64(finalCodeInt)
	}

	fmt.Println(total)
}
