package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var n int
var mapeo map[string]int
var mapeo2 map[int]string
var grafo [][]int
var visitado []bool

// intersect devuelve la intersección de dos slices
func intersect(a, b []int) []int {
	set := make(map[int]bool)
	for _, val := range b {
		set[val] = true
	}
	intersection := []int{}
	for _, val := range a {
		if set[val] {
			intersection = append(intersection, val)
		}
	}
	return intersection
}

// remove elimina un elemento de un slice
func remove(slice []int, elem int) []int {
	for i, v := range slice {
		if v == elem {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

var cliques [][]int

// BronKerbosch encuentra todas las cliques máximas en un grafo
func BronKerbosch(R, P, X []int) {
	//fmt.Println(R, P, X)
	if len(P) == 0 && len(X) == 0 {
		// R es una clique máxima
		cliques = append(cliques, append([]int{}, R...))
		return
	}

	pCopy := append([]int{}, P...)
	for _, v := range pCopy {
		// Expandir R con el nodo v
		newR := append(R, v)
		newP := intersect(P, grafo[v])
		newX := intersect(X, grafo[v])

		// Llamada recursiva
		BronKerbosch(newR, newP, newX)

		// Mover v de P a X
		P = remove(P, v)
		X = append(X, v)
	}
}

func main() {
	//var total uint64 = 0
	tam := 32
	tam = 3380
	n = 0

	reader := bufio.NewReader(os.Stdin)
	grafo = make([][]int, 0)
	mapeo = make(map[string]int)
	mapeo2 = make(map[int]string)
	for c := 0; c < tam; c++ {
		var a, b string
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		vals := strings.Split(line, "-")
		a = vals[0]
		b = vals[1]

		ai, ok := mapeo[a]
		if !ok {
			mapeo[a] = n
			mapeo2[n] = a
			ai = n
			grafo = append(grafo, make([]int, 0))
			n++
		}

		bi, ok := mapeo[b]
		if !ok {
			mapeo[b] = n
			mapeo2[n] = b
			bi = n
			grafo = append(grafo, make([]int, 0))
			n++
		}

		grafo[ai] = append(grafo[ai], bi)
		grafo[bi] = append(grafo[bi], ai)
		//fmt.Printf("%sddd%s\n", a, b)
	}

	cliques = make([][]int, 0)

	R := []int{}
	P := []int{}
	X := []int{}
	for i := 0; i < n; i++ {
		//fmt.Println(mapeo2[i])
		P = append(P, i)
	}

	BronKerbosch(R, P, X)

	maxima := 0
	maxPos := 0
	for i, clique := range cliques {
		//fmt.Println(clique)
		if len(clique) > maxima {
			maxima = len(clique)
			maxPos = i
		}
	}

	cliqueStr := make([]string, 0)
	for _, v := range cliques[maxPos] {
		cliqueStr = append(cliqueStr, mapeo2[v])
	}

	//ordenamos la clique
	sort.Strings(cliqueStr)
	for _, v := range cliqueStr {
		fmt.Println(v)
	}

	fmt.Println(strings.Join(cliqueStr, ","))
}
