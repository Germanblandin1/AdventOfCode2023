// package main

// import (
// 	"fmt"
// 	"sort"
// 	"strings"
// )

// var n int
// var mapeo map[string]int
// var mapeo2 map[int]string
// var grafo [][]int

// // op 0 = 0
// // op 1 = 1
// // op 2 = and
// // op 3 = or
// // op 4 = xor
// var estado []int

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

// func dfs1(nodo int) int {

// 	if estado[nodo] <= 1 {
// 		return estado[nodo]
// 	}
// 	//fmt.Println("nodo", nodo)
// 	a := dfs1(grafo[nodo][0])
// 	b := dfs1(grafo[nodo][1])

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

// func main() {
// 	var total uint64 = 0
// 	tam := 10
// 	tam = 90
// 	tam2 := 36
// 	tam2 = 222
// 	n = 0

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

// 	nodos := make([]string, 0)
// 	for c := 0; c < n; c++ {
// 		if mapeo2[c][0] == 'z' {
// 			nodos = append(nodos, mapeo2[c])
// 		}
// 	}
// 	sort.Strings(nodos)
// 	total = 0
// 	pot := uint64(1)
// 	for _, nodo := range nodos {
// 		val := dfs1(mapeo[nodo])
// 		total += pot * uint64(val)
// 		pot *= 2
// 	}

// 	fmt.Println(total)
// }

// var cambiados []bool

// var solucion []string

// var nocambiar []bool

// var losquecambian []int = []int{0, 0, 0, 0, 0, 0, 0, 0}

// func backtracking(cambios int) bool {
// 	if hayCiclo() {
// 		return false
// 	}
// 	if cambios == 4 {

// 		val := calcular()
// 		//fmt.Printf("%b\n", val)
// 		if val == suma {
// 			for c := 0; c < n; c++ {
// 				if cambiados[c] {
// 					solucion = append(solucion, mapeo2[c])
// 				}
// 			}
// 			return true
// 		}
// 	} else {
// 		for i := subn; i < n; i++ {
// 			if !cambiados[i] {
// 				for j := i + 1; j < n; j++ {
// 					if i == j {
// 						continue
// 					}

// 					if !cambiados[j] {

// 						graOriI0 := grafo[i][0]
// 						graOriI1 := grafo[i][1]
// 						graOriJ0 := grafo[j][0]
// 						graOriJ1 := grafo[j][1]
// 						estadoOriI := estado[i]
// 						estadoOriJ := estado[j]

// 						grafo[i][0] = graOriJ0
// 						grafo[i][1] = graOriJ1
// 						grafo[j][0] = graOriI0
// 						grafo[j][1] = graOriI1
// 						estado[i] = estadoOriJ
// 						estado[j] = estadoOriI

// 						cambiados[i] = true
// 						cambiados[j] = true
// 						losquecambian[cambios*2] = i
// 						losquecambian[cambios*2+1] = j

// 						sepudo := backtracking(cambios + 1)
// 						if sepudo {
// 							return true
// 						}

// 						grafo[i][0] = graOriI0
// 						grafo[i][1] = graOriI1
// 						grafo[j][0] = graOriJ0
// 						grafo[j][1] = graOriJ1
// 						estado[i] = estadoOriI
// 						estado[j] = estadoOriJ

// 						cambiados[i] = false
// 						cambiados[j] = false

// 					}
// 				}
// 			}
// 		}
// 	}

// 	return false

// }

// func marcarNoCambio(nodo int) {
// 	nocambiar[nodo] = false
// 	for i := 0; i < len(grafo[nodo]); i++ {
// 		marcarNoCambio(grafo[nodo][i])
// 	}
// }

// var closure [][]bool

// func warshallClosure() {

// 	// Crear una copia de la matriz para no modificar la original
// 	closure = make([][]bool, n)
// 	for i := 0; i < n; i++ {
// 		closure[i] = make([]bool, n)
// 		for j := 0; j < n; j++ {
// 			closure[i][j] = false
// 		}
// 	}

// 	for i := 0; i < n; i++ {
// 		for j := 0; j < len(grafo[i]); j++ {
// 			closure[i][grafo[i][j]] = true
// 		}
// 	}

// 	// Algoritmo de Warshall
// 	for k := 0; k < n; k++ {
// 		for i := 0; i < n; i++ {
// 			for j := 0; j < n; j++ {
// 				closure[i][j] = closure[i][j] || (closure[i][k] && closure[k][j])
// 			}
// 		}
// 	}
// }