package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var matrix = make([][]rune, 0)
var visited = make([][][]bool, 4)
var contado = make([][]bool, 0)

var movi []int = []int{-1, 0, 1, 0}
var movj []int = []int{0, 1, 0, -1}
var n int
var m int

func dfs(i, j, dir int, pasos *int) {
	//fmt.Printf("%v %v %v %v\n", i, j, dir, *pasos)
	if visited[dir][i][j] {
		return
	}
	if !contado[i][j] {
		*pasos++
		contado[i][j] = true
	}

	visited[dir][i][j] = true
	newi := i + movi[dir]
	newj := j + movj[dir]
	if newi < 0 || newi >= n || newj < 0 || newj >= m {
		return
	}
	if matrix[newi][newj] == '#' {
		dir = (dir + 1) % 4
		newi = i
		newj = j
	}
	dfs(newi, newj, dir, pasos)

}

func main() {
	var total int = 0
	tam := 10
	tam = 130
	reader := bufio.NewReader(os.Stdin)
	inii, inij := 0, 0
	dirini := 0
	for i := 0; i < tam; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		matrix = append(matrix, []rune(line))
		contado = append(contado, make([]bool, len(line)))
		for j := 0; j < len(line); j++ {
			contado[i][j] = false
			if matrix[i][j] == '^' {
				inii = i
				inij = j
				dirini = 0
				matrix[i][j] = '.'
			} else if matrix[i][j] == '>' {
				inii = i
				inij = j
				dirini = 1
				matrix[i][j] = '.'
			} else if matrix[i][j] == 'v' {
				inii = i
				inij = j
				dirini = 2
				matrix[i][j] = '.'
			} else if matrix[i][j] == '<' {
				inii = i
				inij = j
				dirini = 3
				matrix[i][j] = '.'
			}
		}
	}
	n = len(matrix)
	m = len(matrix[0])

	for i := 0; i < 4; i++ {
		visited[i] = make([][]bool, n)
		for j := range visited[i] {
			visited[i][j] = make([]bool, m)
			for k := range visited[i][j] {
				visited[i][j][k] = false
			}
		}
	}

	pasos := 0
	dfs(inii, inij, dirini, &pasos)
	total = pasos

	fmt.Println(total)
}
