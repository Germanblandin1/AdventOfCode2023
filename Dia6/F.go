package main

import "fmt"

var t int
var n, m int

var tiempos []int
var distancias []uint64

func main() {

	n = 1
	var basura string
	fmt.Scan(&basura)
	tiempos = make([]int, n)
	distancias = make([]uint64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tiempos[i])
	}
	fmt.Scan(&basura)
	for i := 0; i < n; i++ {
		fmt.Scan(&distancias[i])
	}
	var total uint64 = 1
	//almenos := false
	for i := 0; i < n; i++ {
		var distancia uint64
		var ganadas uint64 = 0
		for j := 0; j < tiempos[i]; j++ {
			distancia = uint64(j) * uint64(tiempos[i]-j)
			if distancia > distancias[i] {
				ganadas++
				//almenos = true
			}
		}
		total *= ganadas
	}
	fmt.Println(total)

}
