package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var n int
var m int
var mapa [][]rune
var visited [][]bool
var iniI, iniJ int

var movi = []int{0, -1, 0, 1}
var movj = []int{1, 0, -1, 0}
var movc = map[rune]int{'>': 0, '^': 1, '<': 2, 'v': 3}

func empujar(i, j, dir int) bool {

	if i < 0 || i >= n || j < 0 || j >= m || mapa[i][j] == '#' {
		return false
	}

	if mapa[i][j] == '.' {
		return true
	}

	ni := i + movi[dir]
	nj := j + movj[dir]

	esValido := empujar(ni, nj, dir)
	if esValido {
		mapa[ni][nj] = mapa[i][j]
		mapa[i][j] = '.'
	}
	return esValido
}

func PrintMapa() {
	for i := 0; i < n; i++ {
		fmt.Println(string(mapa[i]))
	}
}

func main() {

	var total uint64 = 0
	tam := 10
	tam = 50
	canMov := 20 - tam
	canMov = 70 - tam
	n = tam
	mapa = make([][]rune, n)
	visited = make([][]bool, n)
	reader := bufio.NewReader(os.Stdin)
	for c := 0; c < tam; c++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		mapa[c] = []rune(line)
		m = len(mapa[c])
		visited[c] = make([]bool, m)
		for j := 0; j < m; j++ {
			if mapa[c][j] == '@' {
				iniI = c
				iniJ = j
			}
		}
	}
	var movimientos string = ""
	for i := 0; i < canMov; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		movimientos += line
	}

	//PrintMapa()
	for _, mov := range movimientos {
		//fmt.Println("Movimiento: ", string(mov))
		dir := movc[mov]
		esvalido := empujar(iniI, iniJ, dir)
		if esvalido {
			iniI += movi[dir]
			iniJ += movj[dir]
		}
		//PrintMapa()
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mapa[i][j] == 'O' {
				total += uint64(i*100 + j)
			}
		}
	}

	fmt.Println(total)
}
