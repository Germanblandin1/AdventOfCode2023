package main

import (
	"fmt"
	"io"
	"strconv"
)

var matrixInput [][]rune
var matrixID [][]int
var matrixValue [][]int
var n int
var m int
var index int = 0
var marcaLineal []bool

// movimientos incluyendo diagonal
var movx = []int{0, 0, 1, -1, 1, -1, 1, -1}
var movy = []int{1, -1, 0, 0, 1, -1, -1, 1}

func esUnNumero(c rune) bool {
	return c >= '0' && c <= '9'
}

func marcar(i, j int, numero string) int {

	if j >= m || !esUnNumero(matrixInput[i][j]) {
		val, _ := strconv.Atoi(numero)
		return val
	} else {
		numero += string(matrixInput[i][j])
		value := marcar(i, j+1, numero)
		matrixValue[i][j] = value
		matrixID[i][j] = index
		return value
	}
}

func main() {
	var input string
	matrixInput = make([][]rune, 0)
	for {
		_, err := fmt.Scanln(&input)
		if err == io.EOF {
			break
		}
		matrixInput = append(matrixInput, []rune(input))
	}
	n = len(matrixInput)
	m = len(matrixInput[0])
	matrixID = make([][]int, n)
	matrixValue = make([][]int, n)
	for i := 0; i < n; i++ {
		matrixID[i] = make([]int, m)
		matrixValue[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if esUnNumero(matrixInput[i][j]) && matrixID[i][j] == 0 {
				index++
				marcar(i, j, "")

			}
		}
	}

	//for i := 0; i < n; i++ {
	//	fmt.Println(matrixID[i])
	//}

	var total uint64 = 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrixInput[i][j] == '*' {
				//fmt.Println(i, j, matrixInput[i][j])
				marcaLineal = make([]bool, index+1)
				val1 := -1
				val2 := -1
				valido := true
				for k := 0; k < 8; k++ {
					x := i + movx[k]
					y := j + movy[k]
					if x >= 0 && x < n && y >= 0 && y < m {
						//fmt.Println(x, y, matrixInput[x][y], matrixID[x][y], marcaLineal[matrixID[x][y]])
						if esUnNumero(matrixInput[x][y]) && !marcaLineal[matrixID[x][y]] {
							marcaLineal[matrixID[x][y]] = true
							if val1 == -1 {
								val1 = matrixValue[x][y]
							} else if val2 == -1 {
								val2 = matrixValue[x][y]
							} else {
								valido = false
								break
							}
						}
					}
				}
				if valido && val1 != -1 && val2 != -1 {
					total = total + uint64(val1*val2)
				}
			}
		}
	}
	fmt.Println(total)

}
