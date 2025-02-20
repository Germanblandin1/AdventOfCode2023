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
var mapa2 [][]rune
var iniI, iniJ int

var movi = []int{0, -1, 0, 1}
var movj = []int{1, 0, -1, 0}
var movc = map[rune]int{'>': 0, '^': 1, '<': 2, 'v': 3}

func empujar(i, j, dir int, push bool) bool {

	if i < 0 || i >= n || j < 0 || j >= m || mapa2[i][j] == '#' {
		return false
	}

	if mapa2[i][j] == '.' {
		return true
	}

	ni := i + movi[dir]
	nj := j + movj[dir]

	var val1, val2 bool = true, true
	newJ := 0
	if mapa2[i][j] == '[' && dir%2 == 1 {
		val1 = empujar(ni, nj, dir, push)
		val2 = empujar(ni, nj+1, dir, push)
		newJ = 1
	}

	if mapa2[i][j] == ']' && dir%2 == 1 {
		val1 = empujar(ni, nj, dir, push)
		val2 = empujar(ni, nj-1, dir, push)
		newJ = -1
	}

	if newJ == 0 {
		val1 = empujar(ni, nj, dir, push)
	}

	if push && val1 && val2 && newJ != 0 && mapa2[i][j] != '@' {
		mapa2[ni][nj] = mapa2[i][j]
		mapa2[ni][nj+newJ] = mapa2[i][j+newJ]
		mapa2[i][j] = '.'
		mapa2[i][j+newJ] = '.'
	} else if push && val1 && val2 && (newJ == 0 || mapa2[i][j] == '@') {
		mapa2[ni][nj] = mapa2[i][j]
		mapa2[i][j] = '.'
	}

	return val1 && val2
}

func PrintMapa() {
	for i := 0; i < n; i++ {
		fmt.Println(string(mapa2[i]))
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
	mapa2 = make([][]rune, n)
	reader := bufio.NewReader(os.Stdin)
	for c := 0; c < tam; c++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		mapa[c] = []rune(line)
		m = len(mapa[c])
		mapa2[c] = make([]rune, m*2)
		newM := 0
		for j := 0; j < m; j++ {
			mapa2[c][newM] = mapa[c][j]
			newM++
			mapa2[c][newM] = mapa[c][j]
			newM++
			if mapa[c][j] == '@' {
				iniI = c
				iniJ = newM - 2
				mapa2[c][newM-1] = '.'
			}

			if mapa[c][j] == 'O' {
				mapa2[c][newM-2] = '['
				mapa2[c][newM-1] = ']'
			}
		}
	}
	m = m * 2
	var movimientos string = ""
	for i := 0; i < canMov; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		movimientos += line
	}

	PrintMapa()
	for _, mov := range movimientos {
		fmt.Println("Movimiento: ", string(mov))
		dir := movc[mov]
		esvalido := empujar(iniI, iniJ, dir, false)
		if esvalido {
			empujar(iniI, iniJ, dir, true)
			iniI += movi[dir]
			iniJ += movj[dir]
		}
		PrintMapa()
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mapa2[i][j] == '[' {
				total += uint64(i*100 + j)
			}
		}
	}

	fmt.Println(total)
}
