package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var n int
var m int
var newN int
var newM int
var granjasOrig [][]rune
var granjas [][]rune
var visited [][]bool
var visited2 [][]bool
var visitedBorde [][]int

var movi = []int{0, -1, 0, 1}
var movj = []int{1, 0, -1, 0}

func dfs(i, j int, node rune) int {
	visited[i][j] = true
	area := 1
	//fmt.Printf("i: %d, j: %d, node: %c\n", i, j, node)
	for k := 0; k < 4; k++ {
		ni := i + movi[k]*2
		nj := j + movj[k]*2
		if ni < 0 || ni >= newN || nj < 0 || nj >= newM {
			if k%2 == 0 {
				granjas[i+movi[k]][j+movj[k]] = '|'
			} else {
				granjas[i+movi[k]][j+movj[k]] = '-'
			}
			continue
		} else if granjas[ni][nj] != node {
			if k%2 == 0 {
				granjas[i+movi[k]][j+movj[k]] = '|'
			} else {
				granjas[i+movi[k]][j+movj[k]] = '-'
			}
			continue
		} else if visited[ni][nj] {
			continue
		}

		area += dfs(ni, nj, node)
	}
	return area
}

var bordesCan int

func dfs2(i, j int, node rune) int {
	visited2[i][j] = true
	area := 1
	//fmt.Printf("i: %d, j: %d, node: %c\n", i, j, node)
	for k := 0; k < 4; k++ {
		ni := i + movi[k]*2
		nj := j + movj[k]*2
		if ni < 0 || ni >= newN || nj < 0 || nj >= newM {
			if visitedBorde[i+movi[k]][j+movj[k]] == 0 {
				if k%2 == 0 {
					dfsBorde(i+movi[k], j+movj[k], (k+1)%4, '|', bordesCan)
					dfsBorde(i+movi[k], j+movj[k], (k+3)%4, '|', bordesCan)
				} else {
					dfsBorde(i+movi[k], j+movj[k], (k+1)%4, '-', bordesCan)
					dfsBorde(i+movi[k], j+movj[k], (k+3)%4, '-', bordesCan)
				}
				bordesCan++
			}
			continue
		} else if granjas[ni][nj] != node {
			if visitedBorde[i+movi[k]][j+movj[k]] == 0 {
				if k%2 == 0 {
					dfsBorde(i+movi[k], j+movj[k], (k+1)%4, '|', bordesCan)
					dfsBorde(i+movi[k], j+movj[k], (k+3)%4, '|', bordesCan)
				} else {
					dfsBorde(i+movi[k], j+movj[k], (k+1)%4, '-', bordesCan)
					dfsBorde(i+movi[k], j+movj[k], (k+3)%4, '-', bordesCan)
				}
				bordesCan++
			}
			continue
		} else if visited2[ni][nj] {
			continue
		}

		area += dfs2(ni, nj, node)
	}
	return area
}

func dfsBorde(i, j, dir int, node rune, num int) {
	visitedBorde[i][j] = num

	//fmt.Printf("i: %d, j: %d, node: %c num: %d\n", i, j, node, num)
	ni := i + movi[dir]*2
	nj := j + movj[dir]*2
	//fmt.Printf("ni: %d, nj: %d\n", ni, nj)
	//fmt.Println()
	if ni < 0 || ni >= newN || nj < 0 || nj >= newM {
		return
	} else if !(granjas[ni][nj] == node) {
		return
	} else if visitedBorde[ni][nj] > 0 {
		return
	} else if granjas[i+movi[dir]][j+movj[dir]] == '+' {
		return
	}

	dfsBorde(ni, nj, dir, node, num)

}

func main() {

	var total uint64 = 0
	tam := 10
	tam = 6
	tam = 140
	n = tam
	granjasOrig = make([][]rune, tam)
	reader := bufio.NewReader(os.Stdin)
	for c := 0; c < tam; c++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		granjasOrig[c] = []rune(line)
		m = len(granjasOrig[c])
	}
	newN = n*2 + 1
	newM = m*2 + 1
	granjas = make([][]rune, newN)
	visited = make([][]bool, newN)
	visited2 = make([][]bool, newN)
	visitedBorde = make([][]int, newN)
	for i := 0; i < newN; i++ {
		granjas[i] = make([]rune, newM)
		visited[i] = make([]bool, newM)
		visited2[i] = make([]bool, newM)
		visitedBorde[i] = make([]int, newM)
		for j := 0; j < newM; j++ {
			if i%2 == 0 || j%2 == 0 {
				granjas[i][j] = '.'
			} else {
				granjas[i][j] = granjasOrig[i/2][j/2]
			}
			visited[i][j] = false
			visitedBorde[i][j] = 0
		}

	}
	// for i := 0; i < newN; i++ {
	// 	for j := 0; j < newM; j++ {
	// 		fmt.Printf("%c", granjas[i][j])
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()

	for i := 0; i < newN; i++ {
		for j := 0; j < newM; j++ {
			if !visited[i][j] && granjas[i][j] != '.' && granjas[i][j] != '-' && granjas[i][j] != '|' {
				dfs(i, j, granjas[i][j])

				for i := 0; i < newN; i++ {
					for j := 0; j < newM; j++ {
						count := 0
						if granjas[i][j] == '.' {
							for k := 0; k < 4; k++ {
								ni := i + movi[k]
								nj := j + movj[k]
								if ni < 0 || ni >= newN || nj < 0 || nj >= newM {
									continue
								}
								if granjas[ni][nj] == '-' && (count == 0 || count == 2) {
									count++
								}
								if granjas[ni][nj] == '|' && (count == 1 || count == 0) {
									count += 2
								}
							}
							if count == 3 {
								//fmt.Println("i: ", i, "j: ", j)
								granjas[i][j] = '+'
							}
						}
					}
				}
				bordesCan = 1
				bordesAnt := bordesCan
				a := dfs2(i, j, granjas[i][j])
				fmt.Printf("Rune %c, area %d bordes: %v, bordesAnt %v, bordesDesp %v\n", granjas[i][j], a, bordesCan-bordesAnt, bordesAnt, bordesCan)
				total += uint64(a) * (uint64(bordesCan - bordesAnt))

				for ii := 0; ii < newN; ii++ {
					for jj := 0; jj < newM; jj++ {
						if granjas[ii][jj] == '+' || granjas[ii][jj] == '-' || granjas[ii][jj] == '|' {
							granjas[ii][jj] = '.'
						}

						//fmt.Printf("%d ", visitedBorde[ii][jj])
						visitedBorde[ii][jj] = 0

					}
					//fmt.Println()
				}

			}
		}
	}

	// for i := 0; i < newN; i++ {
	// 	for j := 0; j < newM; j++ {
	// 		fmt.Printf("%d ", visitedBorde[i][j])
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()

	fmt.Println(total)
}
