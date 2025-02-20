package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var n int
var mapeo map[string]int
var mapeo2 map[int]string
var grafo [][]int
var visitado []bool

func main() {
	var total uint64 = 0
	tam := 32
	tam = 3380
	n = 0

	reader := bufio.NewReader(os.Stdin)
	grafo = make([][]int, 0)
	mapeo = make(map[string]int)
	mapeo2 = make(map[int]string)
	for c := 0; c < tam; c++ {
		var a, b string
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		vals := strings.Split(line, "-")
		a = vals[0]
		b = vals[1]

		ai, ok := mapeo[a]
		if !ok {
			mapeo[a] = n
			mapeo2[n] = a
			ai = n
			grafo = append(grafo, make([]int, 0))
			n++
		}

		bi, ok := mapeo[b]
		if !ok {
			mapeo[b] = n
			mapeo2[n] = b
			bi = n
			grafo = append(grafo, make([]int, 0))
			n++
		}

		grafo[ai] = append(grafo[ai], bi)
		grafo[bi] = append(grafo[bi], ai)
		//fmt.Printf("%sddd%s\n", a, b)
	}

	visitado = make([]bool, n)
	//buscamos ciclos de tamaño 3
	total = 0
	for i := 0; i < n; i++ {
		//buscamos una letra t en los nodos
		tieneciclos := false
		//fmt.Println(mapeo2[i])
		if mapeo2[i][0] == 't' {
			for j := 0; j < len(grafo[i]); j++ {
				ni := grafo[i][j]
				if ni == i || visitado[ni] {
					continue
				}
				for k := 0; k < len(grafo[ni]); k++ {
					nni := grafo[ni][k]
					if nni == ni || visitado[nni] {
						continue
					}

					for l := 0; l < len(grafo[nni]); l++ {
						nnni := grafo[nni][l]
						if nnni == nni || visitado[nnni] {
							continue
						}
						if nnni == i {
							fmt.Println(mapeo2[i], mapeo2[ni], mapeo2[nni])
							tieneciclos = true
							total++
							break
						}
					}
				}
			}
			if tieneciclos {
				visitado[i] = true
			}
		}

	}

	fmt.Println(total / 2)
}
