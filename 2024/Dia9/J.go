package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var numbers []int
var n int

func sumatoria(n int) uint64 {
	if n <= 0 {
		return 0
	}
	return uint64(n*(n+1)) / 2
}

func main() {
	var total uint64 = 0
	tam := 1
	//tam = 50
	reader := bufio.NewReader(os.Stdin)

	for c := 0; c < tam; c++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		n = len(line)
		numbers = make([]int, n)
		for i := 0; i < n; i++ {
			numbers[i] = int(line[i] - '0')
		}
	}
	fmt.Println(n)
	if n%2 == 0 {
		n--
	}

	index_izq := 0
	vacios := 0

	index_der := n - 1
	pendientes := 0

	izq := 0

	idIzq := 0
	idDer := n / 2

	result := make([]int, n*9)

	for {

		//fmt.Println("index_izq", index_izq, "index_der", index_der, "izq", izq, "vacios", vacios, "pendientes", pendientes, "idIzq", idIzq, "idDer", idDer, "total", total)
		if vacios == 0 {
			value := numbers[index_izq]
			position := izq + value - 1
			veces := sumatoria(position) - sumatoria(izq-1)
			for i := izq; i <= position; i++ {
				result[i] = idIzq
			}
			total += veces * uint64(idIzq)
			//fmt.Println("-position", position, "veces", veces, "idIzq", idIzq)
			izq = position + 1
			if index_izq+2 == n || index_izq+2 > index_der {
				fmt.Println("idIzq", idIzq, "idDer", idDer)
				break
			}
			vacios = numbers[index_izq+1]
			idIzq++
			index_izq += 2
		} else {
			pendientes = numbers[index_der]
			cuantos := int(math.Min(float64(pendientes), float64(vacios)))
			position := izq + cuantos - 1
			for i := izq; i <= position; i++ {
				result[i] = idDer
			}
			veces := sumatoria(position) - sumatoria(izq-1)
			total += veces * uint64(idDer)
			//fmt.Println("--cuantos", cuantos, "position", position, "veces", veces, "idDer", idDer)
			izq = position + 1
			vacios -= cuantos
			numbers[index_der] -= cuantos
			if numbers[index_der] == 0 {
				index_der -= 2
				idDer--
			}

		}

	}
	for i := 0; i < len(result); i++ {
		fmt.Printf("%v,", result[i])
	}
	fmt.Println()
	fmt.Println(total)
}
