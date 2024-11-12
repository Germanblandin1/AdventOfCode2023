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
var mapa map[string]int
var rmapa map[int]string

var tipos []int
var rtipos []string
var grado []int
var estados []uint64

var memoria [][]bool
var grafo [][]int

func mapear(s string) int {
	//fmt.Println("reciboo", s, len(s))
	if val, ok := mapa[s]; ok {
		//fmt.Println("repite", s)
		return val
	} else {
		//fmt.Println("nuevo", s)
		grafo = append(grafo, make([]int, 0))
		tipos = append(tipos, -1)
		rmapa[actual] = s
		mapa[s] = actual
		actual++
		return actual - 1
	}
}

type Pair struct {
	nodo   int
	pulso  bool
	origen int
}

type Queue []Pair

func (q *Queue) Enqueue(p Pair) {
	*q = append(*q, p)
}

func (q *Queue) Dequeue() (Pair, bool) {
	if len(*q) == 0 {
		return Pair{}, false
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, true
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

var pulsosAltos uint64
var pulsosBajos uint64

func bfs(origenINI int, pulso bool, final int) bool {

	cola := new(Queue)
	cola.Enqueue(Pair{origenINI, pulso, origenINI})

	for !cola.IsEmpty() {
		nodo, _ := cola.Dequeue()
		actual := nodo.nodo
		pulso := nodo.pulso
		origen := nodo.origen

		if tipos[actual] == 0 {
			//broadcast
			for i := 0; i < len(grafo[actual]); i++ {
				cola.Enqueue(Pair{grafo[actual][i], pulso, actual})
				if pulso {
					pulsosAltos++
				} else {
					pulsosBajos++
				}
			}
		} else if tipos[actual] == 1 {
			// flip-flop %
			if !pulso {
				estados[actual] = (uint64(estados[actual]) + 1) % 2
				for i := 0; i < len(grafo[actual]); i++ {
					cola.Enqueue(Pair{grafo[actual][i], estados[actual] == 1, actual})
					if estados[actual] == 1 {
						pulsosAltos++
					} else {
						pulsosBajos++
					}
				}
			}
		} else if tipos[actual] == 2 {
			//conjuncion &
			antes := memoria[actual][origen]
			memoria[actual][origen] = pulso
			if pulso && antes != pulso {
				estados[actual]++
			} else if !pulso && antes != pulso {
				estados[actual]--
			}
			nextPulso := true
			if estados[actual] == uint64(grado[actual]) {
				nextPulso = false
			}
			if actual == final && !nextPulso {
				return true
			}
			for i := 0; i < len(grafo[actual]); i++ {
				cola.Enqueue(Pair{grafo[actual][i], nextPulso, actual})
				if nextPulso {
					pulsosAltos++
				} else {
					pulsosBajos++
				}
			}
		} else if tipos[actual] == 3 {
			if !pulso {
				estados[actual] += 1
			}
		}
	}
	return false
}

var inicio int
var final int

func gcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b uint64) uint64 {
	return abs(a*b) / gcd(a, b)
}

func abs(a uint64) uint64 {
	if a < 0 {
		return -a
	}
	return a
}

func main() {

	fmt.Scan(&t)

	grafo = make([][]int, 0)

	mapa = make(map[string]int)
	tipos = make([]int, 0)
	rmapa = make(map[int]string)
	rtipos = make([]string, 4)

	actual = 0

	for c := 0; c < t; c++ {
		var origenStr string
		fmt.Scan(&origenStr)

		//fmt.Println(origenStr)
		var origen int
		if origenStr[0] == '%' {
			origen = mapear(origenStr[1:])
			tipos[origen] = 1

		} else if origenStr[0] == '&' {
			origen = mapear(origenStr[1:])
			tipos[origen] = 2

		} else if origenStr == "broadcaster" {
			origen = mapear(origenStr)
			tipos[origen] = 0
			inicio = origen
		}

		cuantos := 0
		fmt.Scan(&cuantos)

		for i := 0; i < cuantos; i++ {
			var destinoStr string

			fmt.Scan(&destinoStr)
			destinoStr = strings.Trim(destinoStr, ",")
			//fmt.Println("--->", destinoStr, len(destinoStr))
			destino := mapear(destinoStr)
			if destinoStr == "rx" {
				tipos[destino] = 3
				final = destino
			}
			//fmt.Println(origen, destino)
			grafo[origen] = append(grafo[origen], destino)
		}
	}
	n = actual
	grado = make([]int, n)
	estados = make([]uint64, n)
	memoria = make([][]bool, n)
	for c := 0; c < n; c++ {
		memoria[c] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < len(grafo[i]); j++ {
			if tipos[grafo[i][j]] == 2 {
				grado[grafo[i][j]]++
			}
		}
	}
	pulsosAltos = 0
	pulsosBajos = 0

	rtipos[0] = "broadcaster"
	rtipos[1] = "%"
	rtipos[2] = "&"
	rtipos[3] = "rx"

	importantes := []int{6, 31, 8, 20}
	var resultado uint64 = 1
	for i := 0; i < len(importantes); i++ {
		veces := 0
		for i := 0; i < n; i++ {
			estados[i] = 0
			for j := 0; j < n; j++ {
				memoria[i][j] = false
			}
		}
		for {
			veces++
			val := bfs(inicio, false, importantes[i])
			if val {
				resultado = lcm(resultado, uint64(veces))
				break
			}
			//fmt.Println(veces)
		}
		fmt.Println(veces)

	}
	fmt.Println(resultado)

	//fmt.Println(li, ls, f(li), f(ls))
	// matriz := make([][]uint64, n)
	// for i := 0; i < n; i++ {
	// 	matriz[i] = make([]uint64, n)
	// }
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < len(grafo[i]); j++ {
	// 		matriz[i][grafo[i][j]] = 1
	// 	}
	// }

	// for i := 0; i < n; i++ {
	// 	for j := 0; j < n; j++ {
	// 		if j != n-1 {
	// 			fmt.Print(matriz[i][j], ",")
	// 		} else {
	// 			fmt.Print(matriz[i][j])
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println(mapear("rx") + 1)
	// fmt.Println(mapear("broadcaster") + 1)
	// for i := 0; i < n; i++ {
	// 	if tipos[i] == 2 {
	// 		fmt.Printf("%d ", i+1)
	// 	}
	// }
}
