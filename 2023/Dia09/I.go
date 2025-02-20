package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var basura string
var t int
var n, m int

var secuencias [][]int

func calculateNext(secuencia []int) int {
	//fmt.Println(secuencia)
	subsecuencia := make([]int, len(secuencia)-1)

	zeros := 0

	for i := 0; i < len(secuencia); i++ {
		if secuencia[i] == 0 {
			zeros++
		}
	}

	if zeros == len(secuencia) {
		return 0
	}

	for i := 1; i < len(secuencia); i++ {
		subsecuencia[i-1] = secuencia[i] - secuencia[i-1]
	}
	nexsubsecuencia := calculateNext(subsecuencia)

	return secuencia[0] - nexsubsecuencia
}

func main() {

	n = 3
	n = 200
	//n = 7
	//n = 8
	//m = 2
	//n = 730
	//m = 6

	secuencias = make([][]int, n)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		basura, _ := reader.ReadString('\n')
		basura = strings.Trim(basura, "\n")
		secuenciasStr := strings.Split(basura, " ")
		m = len(secuenciasStr)
		secuencias[i] = make([]int, 0)
		//fmt.Println(secuenciasStr)
		for j := 0; j < m; j++ {
			fmt.Sscanf(secuenciasStr[j], "%d", &t)
			secuencias[i] = append(secuencias[i], t)
		}
		//fmt.Println(secuencias[i])
	}

	suma := 0
	for i := 0; i < n; i++ {
		valor := calculateNext(secuencias[i])
		//fmt.Println(valor)
		suma += valor
	}
	fmt.Println(suma)

}
