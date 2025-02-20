package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var graph = make([][]int, 100)

func initGraph() {
	for i := range graph {
		graph[i] = make([]int, 100)
		for j := range graph[i] {
			graph[i][j] = 0
		}
	}
}

func main() {
	var total int = 0
	tam := 1176
	//tam = 21
	initGraph()
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < tam; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		nums := strings.Split(line, "|")

		A, _ := strconv.Atoi(nums[0])
		B, _ := strconv.Atoi(nums[1])
		fmt.Printf("%v %v %v\n", nums, A, B)
		graph[A][B] = 1
	}

	tam = 1390 - 1177 + 1
	//tam = 6
	for i := 0; i < tam; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		nums := strings.Split(line, ",")
		fmt.Printf("%v\n", nums)

		numeros := make([]int, len(nums))
		for j, num := range nums {
			numeros[j], _ = strconv.Atoi(num)
		}
		esValido := true
		for j := 0; j < len(numeros); j++ {
			for k := j + 1; k < len(numeros); k++ {
				if graph[numeros[k]][numeros[j]] == 1 {
					esValido = false
					aux := numeros[j]
					numeros[j] = numeros[k]
					numeros[k] = aux
				}
			}
		}

		if !esValido {
			medio := numeros[len(numeros)/2]
			total += medio
		}
	}

	fmt.Println(total)
}
