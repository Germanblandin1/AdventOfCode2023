package main

import "fmt"

var basura string
var t int
var n, m int
var INF int = 1000000000

var inii, inij, fini, finj int

var matriz [200][]rune
var dist [200][200]int
var marca [200][200]bool

// >, <, v, ^
var movi []int = []int{0, 0, 1, -1}
var movj []int = []int{1, -1, 0, 0}

//map de runas a enteros
var mapa map[rune]int = map[rune]int{'<': 1, '>': 0, '^': 3, 'v': 2}

func dp(i, j, anti, antj int) int {

	if i == fini && j == finj {
		return 0
	}
	if i < 0 || i >= n || j < 0 || j >= m {
		return -INF
	}
	if matriz[i][j] == '#' {
		return -INF
	}
	//fmt.Println(i, j, anti, antj)

	dist[i][j] = -INF
	mejor := -INF
	for k := 0; k < 4; k++ {

		newi := i + movi[k]
		newj := j + movj[k]

		if newi < 0 || newi >= n || newj < 0 || newj >= m {
			continue
		}
		if matriz[newi][newj] == '#' {
			continue
		}
		if marca[newi][newj] {
			continue
		}
		marca[i][j] = true
		val := dp(newi, newj, i, j) + 1
		marca[i][j] = false
		if val > mejor {
			mejor = val
		}
	}

	return mejor

}

func main() {
	fmt.Scan(&n, &m)
	fmt.Scan(&inii, &inij)
	fmt.Scan(&fini, &finj)
	for i := 0; i < n; i++ {
		fmt.Scan(&basura)
		matriz[i] = []rune(basura)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dist[i][j] = -INF
			marca[i][j] = false
		}
	}
	fmt.Println(dp(inii, inij, -1, -1))

}
