package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var n, m int
var mapa [][]rune
var matrix [][]uint64
var marca [][]bool

var movi []int = []int{0, 0, 1, -1}
var movj []int = []int{1, -1, 0, 0}

func Max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

var contador int

func dp(i, j int) uint64 {

	//fmt.Println(i, j, mapa[i][j]-'0')

	if mapa[i][j] == '9' {
		return 1
	}
	if marca[i][j] {
		return matrix[i][j]
	}

	marca[i][j] = true

	for k := 0; k < 4; k++ {
		newi := i + movi[k]
		newj := j + movj[k]
		if newi < 0 || newi >= n || newj < 0 || newj >= m {
			continue
		}
		if mapa[newi][newj]-mapa[i][j] == 1 {
			matrix[i][j] += dp(newi, newj)
		}
	}
	return matrix[i][j]
}

func main() {
	var total uint64 = 0
	tam := 8
	tam = 53
	reader := bufio.NewReader(os.Stdin)
	mapa = make([][]rune, tam)
	matrix = make([][]uint64, tam)
	marca = make([][]bool, tam)
	n = tam
	for c := 0; c < tam; c++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		m = len(line)
		mapa[c] = []rune(line)
		matrix[c] = make([]uint64, m)
		marca[c] = make([]bool, m)
		for i := 0; i < m; i++ {
			matrix[c][i] = 0
			marca[c][i] = false
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mapa[i][j] == '0' {
				val := dp(i, j)
				//fmt.Println(i, j, val)
				total += val
			}
		}
	}

	fmt.Println(total)
}
