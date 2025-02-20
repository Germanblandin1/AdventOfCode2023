package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var n int
var m int
var inputs [][][]rune
var elemntos [][][]int

func main() {
	var total uint64 = 0
	lineas := 40
	lineas = 4000
	tam := lineas / 8
	n = 7
	m = 5
	inputs = make([][][]rune, tam)
	elemntos = make([][][]int, 2)
	elemntos[0] = make([][]int, 0)
	elemntos[1] = make([][]int, 0)
	reader := bufio.NewReader(os.Stdin)
	for c := 0; c < tam; c++ {
		inputs[c] = make([][]rune, n)
		for i := 0; i < n; i++ {
			line, _ := reader.ReadString('\n')
			line = strings.TrimSpace(line)
			//inputs[c][i] = make([]rune, m)
			inputs[c][i] = []rune(line)
		}

		cual := 0
		//candado
		if inputs[c][0][0] == '#' {
			cual = 0

		} else {
			//llave
			cual = 1
		}
		elemntos[cual] = append(elemntos[cual], make([]int, 5))
		for j := 0; j < m; j++ {
			elemntos[cual][len(elemntos[cual])-1][j] = -1
			for i := 0; i < n; i++ {
				if inputs[c][i][j] == '#' {
					elemntos[cual][len(elemntos[cual])-1][j]++
				}
			}
		}

		reader.ReadString('\n')
	}

	for i := 0; i < len(elemntos[0]); i++ {
		fmt.Println("candado numero", i)
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", elemntos[0][i][j])
		}
		fmt.Println()
	}

	for i := 0; i < len(elemntos[1]); i++ {
		fmt.Println("llave numero", i)
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", elemntos[1][i][j])
		}
		fmt.Println()
	}

	for i := 0; i < len(elemntos[0]); i++ {
		for j := 0; j < len(elemntos[1]); j++ {
			coincide := true
			for k := 0; k < m; k++ {
				//fmt.Println(elemntos[0][i][k], elemntos[1][j][k])
				if elemntos[0][i][k]+elemntos[1][j][k] > 5 {
					coincide = false
				}
			}
			if coincide {
				total++
			}
		}
	}

	fmt.Println(total)
}
