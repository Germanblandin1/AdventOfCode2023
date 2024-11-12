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

var dp [][][]int
var marca [][][]bool

func solve(i, j, k int) int {
	if j == m {
		fmt.Println(i, j, k)
		return 1
	}

	if i >= n || j > m {
		return 0
	}
	if k > grupos[actual][j] {
		return 0
	}
	if marca[i][j][k] {
		return dp[i][j][k]
	}
	//marca[i][j][k] = true
	dp[i][j][k] = 0
	//???.### 1,1,3
	if lineas[actual][i] == '.' {
		if k == 0 {
			dp[i][j][k] = solve(i+1, j, k)
		} else if k < grupos[actual][j] {
			return 0
		} else if k == grupos[actual][j] {
			dp[i][j][k] = solve(i+1, j+1, 0)
		}
	} else if lineas[actual][i] == '#' {
		dp[i][j][k] = solve(i+1, j, k+1)
	} else if lineas[actual][i] == '?' {
		//imagina que es un #
		dp[i][j][k] = solve(i+1, j, k+1)
		//imagina que es un .
		if k == 0 {
			dp[i][j][k] += solve(i+1, j, k)
		} else if k < grupos[actual][j] {
			dp[i][j][k] += 0
		} else if k == grupos[actual][j] {
			dp[i][j][k] += solve(i+1, j+1, 0)
		}
	}
	return dp[i][j][k]
}

func main() {

	fmt.Scan(&t)
	lineas = make([][]rune, t)
	grupos = make([][]int, t)
	for c := 0; c < t; c++ {
		fmt.Scan(&basura)
		basura = basura + "."
		lineas[c] = make([]rune, len(basura))
		lineas[c] = []rune(basura)
		fmt.Scan(&basura)
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
		dp = make([][][]int, n+1)
		marca = make([][][]bool, n+1)
		for i := 0; i < n+1; i++ {
			dp[i] = make([][]int, m+1)
			marca[i] = make([][]bool, m+1)
			for j := 0; j < m+1; j++ {
				dp[i][j] = make([]int, n+1)
				marca[i][j] = make([]bool, n+1)
				for k := 0; k < n+1; k++ {
					dp[i][j][k] = -1
					marca[i][j][k] = false
				}
			}
		}
		actual = c
		valor := solve(0, 0, 0)
		fmt.Println(string(lineas[actual]), grupos[actual], valor)
		total += valor
	}
	fmt.Println(total)

}
