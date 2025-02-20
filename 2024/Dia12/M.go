package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var n int
var m int
var granjas [][]rune
var visited [][]bool

var movi = []int{0, -1, 0, 1}
var movj = []int{1, 0, -1, 0}

func dfs(i, j int, node rune) (int, int) {
	visited[i][j] = true
	area := 1
	perimeter := 0

	for k := 0; k < 4; k++ {
		ni := i + movi[k]
		nj := j + movj[k]
		if ni < 0 || ni >= n || nj < 0 || nj >= m {
			perimeter++
			continue
		} else if granjas[ni][nj] != node {
			perimeter++
			continue
		} else if visited[ni][nj] {
			continue
		}
		a, p := dfs(ni, nj, node)
		area += a
		perimeter += p
	}
	return area, perimeter
}

func main() {

	var total uint64 = 0
	tam := 10
	tam = 140
	n = tam
	granjas = make([][]rune, n)
	visited = make([][]bool, n)
	reader := bufio.NewReader(os.Stdin)
	for c := 0; c < tam; c++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		granjas[c] = []rune(line)
		m = len(granjas[c])
		visited[c] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !visited[i][j] {
				a, p := dfs(i, j, granjas[i][j])
				total += uint64(a * p)
			}
		}
	}

	fmt.Println(total)
}
