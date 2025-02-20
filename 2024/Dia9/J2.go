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

var numbers []int
var posiciones []int
var n int

func sumatoria(n int) uint64 {
	if n <= 0 {
		return 0
	}
	return uint64(n*(n+1)) / 2
}

var agrupados []PriorityQueue[int]
var marca []bool

func main() {
	var total uint64 = 0
	tam := 1
	//tam = 50
	reader := bufio.NewReader(os.Stdin)
	agrupados = make([]PriorityQueue[int], 10)

	for i := 0; i < 10; i++ {
		agrupados[i] = PriorityQueue[int]{}
	}
	for c := 0; c < tam; c++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		n = len(line)
		marca = make([]bool, n*9/2+1)
		numbers = make([]int, n)
		posiciones = make([]int, n)
		pos := 0
		ant := 0
		for i := 0; i < n; i++ {
			numbers[i] = int(line[i] - '0')
			pos += ant
			posiciones[i] = pos
			if i%2 == 1 {
				agrupados[numbers[i]].Insert(pos, pos)
			}

			ant = numbers[i]

		}
	}

	fmt.Println(n)

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("i", i, "tam", agrupados[i].Len(), "files", agrupados[i].String())
	// }
	if n%2 == 0 {
		n--
	}

	index_der := n - 1

	pendientes := 0

	idDer := n / 2

	result := make([]int, n*9)

	for {
		fmt.Println("numbers[index_der]", numbers[index_der], "idDer", idDer, "total", total)
		pendientes = -1
		minID := 100000000
		posMin := 10000000
		value := numbers[index_der]
		for i := value; i <= 9; i++ {
			if agrupados[i].Len() == 0 {
				continue
			}
			primero, _ := agrupados[i].Peek()
			if primero < minID && primero < posiciones[index_der] {
				minID = primero
				posMin = i
			}
		}

		fmt.Println("minID", minID, "posMin", posMin)
		tizq := -1
		pendientes = value
		if posMin != 10000000 {
			//fmt.Println(agrupados[posMin].String())
			agrupados[posMin].Extract()
			tizq = minID
			if posMin > value {
				fmt.Println("posMin", posMin, "value", value, "minID", minID, posMin-value, minID+value)
				agrupados[posMin-value].Insert(minID+value, minID+value)
			}
		} else {
			tizq = posiciones[index_der]
		}

		//fmt.Println("pendientes", pendientes)
		cuantos := pendientes
		position := tizq + cuantos - 1
		for i := tizq; i <= position; i++ {
			result[i] = idDer
		}
		veces := sumatoria(position) - sumatoria(tizq-1)
		total += veces * uint64(idDer)
		fmt.Println("--cuantos", cuantos, "tizq", tizq, "position", position, "veces", veces, "idDer", idDer)

		index_der -= 2
		idDer--
		if index_der < 0 {
			break
		}

	}
	for i := 0; i < len(result); i++ {
		fmt.Printf("%v,", result[i])
	}
	fmt.Println()
	fmt.Println(total)
}
