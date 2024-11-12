package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var basura string
var t int
var n, m int
var INF int = 1000000000

var grafo [][]int
var mapa map[string]int

func globalMinCut(mat [][]int) (int, []int) {
	best := math.MaxInt32
	var bestCo []int
	n := len(mat)
	co := make([][]int, n)

	for i := 0; i < n; i++ {
		co[i] = []int{i}
	}

	for ph := 1; ph < n; ph++ {
		w := make([]int, len(mat[0]))
		copy(w, mat[0])
		var s, t int
		for it := 0; it < n-ph; it++ {
			w[t] = math.MinInt32
			s, t = t, maxElementIndex(w)
			for i := 0; i < n; i++ {
				w[i] += mat[t][i]
			}
		}
		if w[t]-mat[t][t] < best {
			best = w[t] - mat[t][t]
			bestCo = make([]int, len(co[t]))
			copy(bestCo, co[t])
		}
		co[s] = append(co[s], co[t]...)
		for i := 0; i < n; i++ {
			mat[s][i] += mat[t][i]
		}
		for i := 0; i < n; i++ {
			mat[i][s] = mat[s][i]
		}
		mat[0][t] = math.MinInt32
	}

	return best, bestCo
}

func maxElementIndex(arr []int) int {
	maxIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

func mapear(s string) int {
	if val, ok := mapa[s]; ok {
		return val
	} else {
		//fmt.Println("mapeando", s, "blbla")
		mapa[s] = n
		grafo = append(grafo, make([]int, 0))
		n++
		return mapa[s]
	}
}

func main() {
	fmt.Scan(&t)
	reader := bufio.NewReader(os.Stdin)
	basura, _ = reader.ReadString('\n')
	mapa = make(map[string]int)
	grafo = make([][]int, 0)

	n = 0
	for i := 0; i < t; i++ {
		//leer linea entera
		var basura string = ""
		basura, _ = reader.ReadString('\n')
		//basura, _ = reader.ReadString('\n')
		basura = strings.Trim(basura, "\n")
		basura = strings.TrimSpace(basura)
		//quitar los \n
		//separar por espacios
		lista := strings.Split(basura, " ")
		//fmt.Println(lista)
		//fmt.Println(len(lista))
		//fmt.Println(lista[len(lista)-1])
		origen := mapear(lista[0])
		for j := 1; j < len(lista); j++ {
			var destinostr string

			destinostr = lista[j]

			destino := mapear(destinostr)
			grafo[origen] = append(grafo[origen], destino)
			grafo[destino] = append(grafo[destino], origen)
		}
	}
	fmt.Println(n)
	//crear matriz de adyacencia
	matriz := make([][]int, n)
	for i := 0; i < n; i++ {
		matriz[i] = make([]int, n)
		for j := 0; j < len(grafo[i]); j++ {
			matriz[i][grafo[i][j]] = 1
		}
	}

	mincut, cut := globalMinCut(matriz)
	fmt.Println(mincut)
	fmt.Println(len(cut) * (n - len(cut)))

}
