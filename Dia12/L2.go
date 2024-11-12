package main

import (
	"fmt"
	"strings"
)

var basura string
var t int
var n, m int
var INF int = 1000000000
var actual int

var lineas [][]rune
var grupos [][]int

var dp [][]int
var marca [][]bool

func solve(i, j int) int {
	if j == m {
		for pos := i; pos < n; pos++ {
			if lineas[actual][pos] == '#' {
				return 0
			}
		}
		//fmt.Println(string(lineas[actual]))
		return 1
	}

	if i >= n || j >= m {
		return 0
	}
	if marca[i][j] {
		//fmt.Println("hola")
		return dp[i][j]
	}
	//marca[i][j] = true
	dp[i][j] = 0

	if lineas[actual][i] == '#' || lineas[actual][i] == '?' {
		pos := i
		agarrar := grupos[actual][j]
		llevo := 0
		for {
			if pos >= n || llevo > agarrar {
				break
			}
			if lineas[actual][pos] == '.' {
				if llevo == agarrar {
					copia := make([]rune, n)
					copy(copia, lineas[actual])
					for k := 0; k < agarrar; k++ {
						lineas[actual][i+k] = '#'
					}
					lineas[actual][pos] = '.'
					dp[i][j] += solve(pos+1, j+1)
					copy(lineas[actual], copia)
				}
				break
			} else if lineas[actual][pos] == '#' {
				llevo++
			} else if lineas[actual][pos] == '?' {
				if llevo == agarrar {
					copia := make([]rune, n)
					copy(copia, lineas[actual])
					for k := 0; k < agarrar; k++ {
						lineas[actual][i+k] = '#'
					}
					lineas[actual][pos] = '.'
					dp[i][j] += solve(pos+1, j+1)
					copy(lineas[actual], copia)
				}
				llevo++
			}
			pos++
		}
	} else if lineas[actual][i] == '.' {
		dp[i][j] += solve(i+1, j)
	}
	if lineas[actual][i] == '?' {
		lineas[actual][i] = '.'
		dp[i][j] += solve(i+1, j)
		lineas[actual][i] = '?'
	}

	return dp[i][j]
}

func main() {

	fmt.Scan(&t)
	lineas = make([][]rune, t)
	grupos = make([][]int, t)
	for c := 0; c < t; c++ {
		fmt.Scan(&basura)
		basura = basura + "?" + basura + "?" + basura + "?" + basura + "?" + basura
		basura = basura + "."
		lineas[c] = make([]rune, len(basura))
		lineas[c] = []rune(basura)
		fmt.Scan(&basura)
		basura = basura + "," + basura + "," + basura + "," + basura + "," + basura
		enteros := strings.Split(basura, ",")
		grupos[c] = make([]int, len(enteros))
		for i := 0; i < len(enteros); i++ {
			fmt.Sscan(enteros[i], &grupos[c][i])
		}
		//fmt.Println(string(lineas[c]))
		//fmt.Println(grupos[c])
	}

	total := 0
	for c := 0; c < t; c++ {
		n = len(lineas[c])
		m = len(grupos[c])
		dp = make([][]int, n+1)
		marca = make([][]bool, n+1)
		for i := 0; i < n+1; i++ {
			dp[i] = make([]int, m+1)
			marca[i] = make([]bool, m+1)
			for j := 0; j < m+1; j++ {

				dp[i][j] = -INF
				marca[i][j] = false

			}
		}
		actual = c
		valor := solve(0, 0)
		fmt.Println(string(lineas[actual]), grupos[actual], valor)
		total += valor
	}
	fmt.Println(total)

}
