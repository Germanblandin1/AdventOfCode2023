// package main

// import (
// 	"container/heap"
// 	"fmt"
// 	"sort"
// 	"strings"
// )

// // Item representa un elemento en la cola de prioridad.
// type Item[T any] struct {
// 	Value    T   // El valor almacenado en el elemento.
// 	Priority int // La prioridad del elemento. Un número menor indica mayor prioridad.
// }

// // PriorityQueue es una cola de prioridad genérica basada en un heap.
// type PriorityQueue[T any] []*Item[T]

// // Len devuelve la cantidad de elementos en la cola.
// func (pq PriorityQueue[T]) Len() int { return len(pq) }

// // Less compara las prioridades de dos elementos.
// // Nota: Cambia la comparación si prefieres una cola de prioridad máxima.
// func (pq PriorityQueue[T]) Less(i, j int) bool {
// 	return pq[i].Priority < pq[j].Priority // Menor prioridad tiene más prioridad.
// }

// // Swap intercambia dos elementos en la cola.
// func (pq PriorityQueue[T]) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// }

// // Push agrega un elemento a la cola. Este método es utilizado por `container/heap`.
// func (pq *PriorityQueue[T]) Push(x any) {
// 	item := x.(*Item[T])
// 	*pq = append(*pq, item)
// }

// // Pop elimina y devuelve el elemento con mayor prioridad. Este método es utilizado por `container/heap`.
// func (pq *PriorityQueue[T]) Pop() any {
// 	old := *pq
// 	n := len(old)
// 	item := old[n-1]
// 	*pq = old[0 : n-1]
// 	return item
// }

// // Insert agrega un elemento a la cola de prioridad.
// func (pq *PriorityQueue[T]) Insert(value T, priority int) {
// 	heap.Push(pq, &Item[T]{Value: value, Priority: priority})
// }

// // Extract devuelve el elemento con mayor prioridad y lo elimina de la cola.
// // Si la cola está vacía, devuelve el valor por defecto de `T` y `false`.
// func (pq *PriorityQueue[T]) Extract() (T, bool) {
// 	if pq.Len() == 0 {
// 		var zero T
// 		return zero, false
// 	}
// 	item := heap.Pop(pq).(*Item[T])
// 	return item.Value, true
// }

// // Peek devuelve el elemento con mayor prioridad sin eliminarlo.
// // Si la cola está vacía, devuelve el valor por defecto de `T` y `false`.
// func (pq *PriorityQueue[T]) Peek() (T, bool) {
// 	if pq.Len() == 0 {
// 		var zero T
// 		return zero, false
// 	}
// 	return (*pq)[0].Value, true
// }

// // String devuelve un string con todos los elementos de la cola de prioridad.
// func (pq *PriorityQueue[T]) String() string {
// 	result := "PriorityQueue:\n"
// 	for _, item := range *pq {
// 		result += fmt.Sprintf("Value: %v, Priority: %d\n", item.Value, item.Priority)
// 	}
// 	return result
// }

// var n int
// var subn int
// var mapeo map[string]int
// var mapeo2 map[int]string
// var grafo [][]int

// // op 0 = 0
// // op 1 = 1
// // op 2 = and
// // op 3 = or
// // op 4 = xor
// var estado []int
// var estadoOriginal []int

// func addNodo(a string) int {
// 	ai, ok := mapeo[a]
// 	if !ok {
// 		mapeo[a] = n
// 		mapeo2[n] = a
// 		ai = n
// 		grafo = append(grafo, make([]int, 0))
// 		estado = append(estado, 0)
// 		n++
// 	}
// 	return ai
// }

// func dfs(nodo int, visitado []bool) int {
// 	//fmt.Println("dfs ", mapeo2[nodo])
// 	if estado[nodo] <= 1 {
// 		return estado[nodo]
// 	}

// 	if visitado[nodo] {
// 		fmt.Println("empezo ", mapeo2[nodo])
// 		return -1
// 	}
// 	visitado[nodo] = true
// 	a := dfs(grafo[nodo][0], visitado)
// 	b := dfs(grafo[nodo][1], visitado)

// 	if a == -1 {
// 		fmt.Println("a ", mapeo2[nodo])
// 		return -1
// 	}

// 	if b == -1 {
// 		fmt.Println("b ", mapeo2[nodo])
// 		return -1
// 	}

// 	newEstado := 0
// 	if estado[nodo] == 2 {
// 		newEstado = a & b
// 	}
// 	if estado[nodo] == 3 {
// 		newEstado = a | b
// 	}
// 	if estado[nodo] == 4 {
// 		newEstado = a ^ b
// 	}
// 	estado[nodo] = newEstado
// 	return newEstado
// }

// var nodos []string

// var aristasOriginales [][][]int
// var estadosOrig [][]int

// func cambiarAristas(cambios int, cambiados [][]int) {
// 	for c := 0; c < cambios; c++ {
// 		origen := cambiados[c][0]
// 		destino := cambiados[c][1]
// 		//fmt.Println("cambios ", cambios)
// 		//fmt.Println("cambiando ", mapeo2[origen], mapeo2[destino])
// 		aristasOriginales[c][0][0] = grafo[origen][0]
// 		aristasOriginales[c][0][1] = grafo[origen][1]
// 		aristasOriginales[c][1][0] = grafo[destino][0]
// 		aristasOriginales[c][1][1] = grafo[destino][1]
// 		estadosOrig[c][0] = estado[origen]
// 		estadosOrig[c][1] = estado[destino]

// 		grafo[origen][0] = aristasOriginales[c][1][0]
// 		grafo[origen][1] = aristasOriginales[c][1][1]
// 		grafo[destino][0] = aristasOriginales[c][0][0]
// 		grafo[destino][1] = aristasOriginales[c][0][1]
// 		estado[origen] = estadosOrig[c][1]
// 		estado[destino] = estadosOrig[c][0]
// 	}
// }

// func restaurarAristas(cambios int, cambiados [][]int) {
// 	for c := 0; c < cambios; c++ {
// 		origen := cambiados[c][0]
// 		destino := cambiados[c][1]

// 		grafo[origen][0] = aristasOriginales[c][0][0]
// 		grafo[origen][1] = aristasOriginales[c][0][1]
// 		grafo[destino][0] = aristasOriginales[c][1][0]
// 		grafo[destino][1] = aristasOriginales[c][1][1]
// 		estado[origen] = estadosOrig[c][0]
// 		estado[destino] = estadosOrig[c][1]
// 	}
// }

// func modificarX(val uint64) {

// 	//guarda bit a bit el valor dado en los estados de x
// 	for c := 0; c < n; c++ {
// 		if mapeo2[c][0] == 'x' {
// 			estadoOriginal[c] = int(val % 2)
// 			val /= 2
// 		}
// 	}
// }

// func modificarY(val uint64) {

// 	//guarda bit a bit el valor dado en los estados de x
// 	for c := 0; c < n; c++ {
// 		if mapeo2[c][0] == 'y' {
// 			estadoOriginal[c] = int(val % 2)
// 			val /= 2
// 		}
// 	}
// }

// func calcularReal(cambios int, cambiados [][]int) (uint64, uint64, uint64, uint64) {
// 	a := uint64(1<<45 - 1)
// 	b := uint64(1 << 45)
// 	modificarX(a)
// 	modificarY(b)
// 	suma1 := a + b
// 	val1 := calcular(cambios, cambiados)
// 	//return suma1, val1, 0, 0
// 	c := uint64(0)
// 	d := uint64(0)
// 	modificarX(c)
// 	modificarY(d)
// 	suma2 := c + d
// 	val2 := calcular(cambios, cambiados)
// 	return suma1, val1, suma2, val2
// }

// func calcular(cambios int, cambiados [][]int) uint64 {
// 	copy(estado, estadoOriginal)
// 	cambiarAristas(cambios, cambiados)

// 	if hayCiclo() {
// 		restaurarAristas(cambios, cambiados)
// 		return 0
// 	}
// 	visitado := make([]bool, n)
// 	var total uint64 = 0
// 	pot := uint64(1)
// 	for _, nodo := range nodos {
// 		//fmt.Println("nodo ", nodo)
// 		val := dfs(mapeo[nodo], visitado)
// 		if val == -1 {
// 			fmt.Println("termino ", nodo)
// 			return 0
// 		}
// 		total += pot * uint64(val)
// 		pot *= 2
// 	}

// 	restaurarAristas(cambios, cambiados)

// 	return total
// }

// func diffBits2(a, b uint64) int {
// 	c := a ^ b
// 	count := 0
// 	for c > 0 {
// 		if c&1 == 1 {
// 			count++
// 		}
// 		c >>= 1
// 	}
// 	return count
// }

// func diffBits(a, b, c, d uint64) int {
// 	return diffBits2(a, b) + diffBits2(c, d)
// }

// func dfsCiclo(nodo int, visitado, apilado []bool) bool {
// 	visitado[nodo] = true
// 	apilado[nodo] = true
// 	for i := 0; i < len(grafo[nodo]); i++ {
// 		if !visitado[grafo[nodo][i]] {
// 			if dfsCiclo(grafo[nodo][i], visitado, apilado) {
// 				return true
// 			}
// 		} else if apilado[grafo[nodo][i]] {
// 			return true
// 		}
// 	}
// 	apilado[nodo] = false
// 	return false

// }

// func hayCiclo() bool {
// 	visitado := make([]bool, n)
// 	apilado := make([]bool, n)
// 	for i := 0; i < n; i++ {
// 		if !visitado[i] {
// 			if dfsCiclo(i, visitado, apilado) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// var suma uint64

// type Estado struct {
// 	cambios   int
// 	cambiados [][]int
// }

// func estaCambiado(nodo, cambios int, cambiados [][]int) bool {
// 	for c := 0; c < cambios; c++ {
// 		if cambiados[c][0] == nodo || cambiados[c][1] == nodo {
// 			return true
// 		}
// 	}
// 	return false
// }

// var ordenados []string

// func AStar() []string {
// 	pq := make(PriorityQueue[Estado], 0)
// 	estadoInicial := Estado{cambios: 0, cambiados: make([][]int, 0)}
// 	pq.Insert(estadoInicial, diffBits(calcularReal(estadoInicial.cambios, estadoInicial.cambiados)))

// 	for pq.Len() > 0 {
// 		estadoActual, _ := pq.Extract()
// 		if estadoActual.cambios == 4 {
// 			if a, b, c, d := calcularReal(estadoActual.cambios, estadoActual.cambiados); a == b && c == d {
// 				fmt.Printf("encontrado %b\n %b\n %b\n %b\n\n", a, b, c, d)
// 				fmt.Println("encontrado")
// 				solucion := make([]string, 0)

// 				for c := 0; c < estadoActual.cambios; c++ {
// 					solucion = append(solucion, mapeo2[estadoActual.cambiados[c][0]])
// 					solucion = append(solucion, mapeo2[estadoActual.cambiados[c][1]])
// 					//fmt.Println("cambie ", mapeo2[estadoActual.cambiados[c][0]], mapeo2[estadoActual.cambiados[c][1]])
// 				}
// 				return solucion

// 			}
// 		} else {
// 			for o1 := 0; o1 < len(ordenados); o1++ {
// 				i := mapeo[ordenados[o1]]
// 				if estaCambiado(i, estadoActual.cambios, estadoActual.cambiados) {
// 					continue
// 				}
// 				for o2 := o1 + 1; o2 < len(ordenados); o2++ {
// 					j := mapeo[ordenados[o2]]
// 					if i == j {
// 						continue
// 					}
// 					if estaCambiado(j, estadoActual.cambios, estadoActual.cambiados) {
// 						continue
// 					}

// 					estadoSiguiente := Estado{cambios: estadoActual.cambios + 1, cambiados: make([][]int, 0)}
// 					for c := 0; c < estadoActual.cambios; c++ {
// 						estadoSiguiente.cambiados = append(estadoSiguiente.cambiados, estadoActual.cambiados[c])
// 					}
// 					estadoSiguiente.cambiados = append(estadoSiguiente.cambiados, []int{i, j})
// 					pq.Insert(estadoSiguiente, diffBits(calcularReal(estadoSiguiente.cambios, estadoSiguiente.cambiados)))

// 				}
// 			}
// 		}

// 	}

// 	return nil
// }

// func main() {
// 	var total uint64 = 0
// 	tam := 10
// 	tam = 12
// 	tam = 90
// 	tam2 := 36
// 	tam2 = 6
// 	tam2 = 222
// 	n = 0

// 	subn = tam
// 	//reader := bufio.NewReader(os.Stdin)
// 	grafo = make([][]int, 0)
// 	mapeo = make(map[string]int)
// 	mapeo2 = make(map[int]string)
// 	estado = make([]int, 0)

// 	for c := 0; c < tam; c++ {
// 		var a string
// 		var num int
// 		fmt.Scanf("%s %d\n", &a, &num)
// 		a = strings.Trim(a, " ")
// 		a = strings.Trim(a, ":")

// 		ai := addNodo(a)
// 		estado[ai] = num
// 	}

// 	for c := 0; c < tam2; c++ {
// 		var a, b, op, c string
// 		fmt.Scanf("%s %s %s -> %s\n", &a, &op, &b, &c)
// 		a1 := addNodo(a)
// 		b1 := addNodo(b)
// 		c1 := addNodo(c)

// 		opInt := 0
// 		if op == "AND" {
// 			opInt = 2
// 		}
// 		if op == "OR" {
// 			opInt = 3
// 		}
// 		if op == "XOR" {
// 			opInt = 4
// 		}
// 		grafo[c1] = append(grafo[c1], a1)
// 		grafo[c1] = append(grafo[c1], b1)
// 		estado[c1] = opInt
// 	}

// 	//calcular x
// 	valX := uint64(0)
// 	pot := uint64(1)
// 	for c := 0; c < n; c++ {
// 		if mapeo2[c][0] == 'x' {
// 			valX += pot * uint64(estado[c])
// 			pot *= 2
// 		}
// 	}

// 	valY := uint64(0)
// 	pot = 1
// 	for c := 0; c < n; c++ {
// 		if mapeo2[c][0] == 'y' {
// 			valY += pot * uint64(estado[c])
// 			pot *= 2
// 		}
// 	}
// 	estadoOriginal = make([]int, n)
// 	for c := 0; c < n; c++ {
// 		estadoOriginal[c] = estado[c]
// 	}

// 	suma = valX + valY
// 	fmt.Printf("valX %b\n", valX)
// 	fmt.Printf("valY %b\n", valY)
// 	fmt.Printf("suma %b\n", suma)

// 	nodos = make([]string, 0)
// 	for c := 0; c < n; c++ {
// 		if mapeo2[c][0] == 'z' {
// 			nodos = append(nodos, mapeo2[c])
// 		}
// 	}
// 	sort.Strings(nodos)

// 	presol := calcular(0, make([][]int, 0))
// 	fmt.Printf("pres %b\n", presol)

// 	aristasOriginales = make([][][]int, 4)
// 	estadosOrig = make([][]int, 4)
// 	for c := 0; c < 4; c++ {
// 		aristasOriginales[c] = make([][]int, 2)
// 		aristasOriginales[c][0] = make([]int, 2)
// 		aristasOriginales[c][1] = make([]int, 2)
// 		estadosOrig[c] = make([]int, 2)
// 	}

// 	ordenados = make([]string, 0)
// 	for c := subn; c < n; c++ {
// 		ordenados = append(ordenados, mapeo2[c])
// 	}
// 	sort.Strings(ordenados)

// 	solucion := AStar()

// 	sort.Strings(solucion)

// 	fmt.Println(strings.Join(solucion, ","))

// 	fmt.Println(total)

// }
