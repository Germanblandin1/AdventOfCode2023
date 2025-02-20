package main

import "fmt"

var basura string
var t int
var n, m int
var INF int = 1000000000
var actual int

var matriz [][]rune

//0 hacia derecha, 1 hacia izquierda, 2 hacia abajo, 3 hacia arriba
var movi []int = []int{0, 0, 1, -1}
var movj []int = []int{1, -1, 0, 0}

var visitado [][][]bool
var marca [][]bool

//caracteres .,-,|, \ y /

func next(i, j, dir int) []int {
	nextDirs := make([]int, 0)
	if matriz[i][j] == '.' {
		nextDirs = append(nextDirs, dir)
	} else if matriz[i][j] == '-' {
		if dir == 0 || dir == 1 {
			nextDirs = append(nextDirs, dir)
		}
		if dir == 2 || dir == 3 {
			nextDirs = append(nextDirs, 0)
			nextDirs = append(nextDirs, 1)
		}
	} else if matriz[i][j] == '|' {
		if dir == 2 || dir == 3 {
			nextDirs = append(nextDirs, dir)
		}
		if dir == 0 || dir == 1 {
			nextDirs = append(nextDirs, 2)
			nextDirs = append(nextDirs, 3)
		}
	} else if matriz[i][j] == '\\' {
		if dir == 0 {
			//desde derecha baja
			nextDirs = append(nextDirs, 2)
		}
		if dir == 1 {
			//desde izquierda sube
			nextDirs = append(nextDirs, 3)
		}
		if dir == 2 {
			//desde abajo izquierda
			nextDirs = append(nextDirs, 0)
		}
		if dir == 3 {
			//desde arriba derecha
			nextDirs = append(nextDirs, 1)
		}
	} else if matriz[i][j] == '/' {
		if dir == 0 {
			//desde derecha sube
			nextDirs = append(nextDirs, 3)
		}
		if dir == 1 {
			//desde izquierda baja
			nextDirs = append(nextDirs, 2)
		}
		if dir == 2 {
			//desde abajo derecha
			nextDirs = append(nextDirs, 1)
		}
		if dir == 3 {
			//desde arriba izquierda
			nextDirs = append(nextDirs, 0)
		}
	}
	return nextDirs

}

var total int

func dfs(i, j, dir int) {
	if !marca[i][j] {
		marca[i][j] = true
		total++
	}
	visitado[i][j][dir] = true

	newi := i + movi[dir]
	newj := j + movj[dir]
	if newi >= 0 && newi < n && newj >= 0 && newj < m {
		nextDirs := next(newi, newj, dir)
		for _, nextDir := range nextDirs {
			if !visitado[newi][newj][nextDir] {
				dfs(newi, newj, nextDir)
			}
		}
	}

}

func limpiar() {
	visitado = make([][][]bool, n)
	marca = make([][]bool, n)
	for i := 0; i < n; i++ {
		visitado[i] = make([][]bool, m)
		marca[i] = make([]bool, m)
		for j := 0; j < m; j++ {
			visitado[i][j] = make([]bool, 4)
		}
	}

}

var maximo int

func solve(i, j, dir int) {
	limpiar()
	total = 0
	nextDirs := next(i, j, dir)
	for _, nextDir := range nextDirs {
		if !visitado[i][j][nextDir] {
			dfs(i, j, nextDir)
		}
	}
	if total > maximo {
		maximo = total
	}
}

func main() {
	fmt.Scan(&n, &m)
	matriz = make([][]rune, 0)
	visitado = make([][][]bool, n)
	marca = make([][]bool, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&basura)
		matriz = append(matriz, []rune(basura))
		visitado[i] = make([][]bool, m)
		marca[i] = make([]bool, m)
		for j := 0; j < m; j++ {
			visitado[i][j] = make([]bool, 4)
		}
	}

	maximo = 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 || i == n-1 || j == m-1 {
				if i == 0 {
					solve(i, j, 2)
				}
				if j == 0 {
					solve(i, j, 0)
				}
				if i == n-1 {
					solve(i, j, 3)
				}
				if j == m-1 {
					solve(i, j, 1)
				}
			}
		}
	}

	fmt.Println(maximo)

}
