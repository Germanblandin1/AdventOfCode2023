package main

import (
	"fmt"
)

var t int
var n, m int

type conversion struct {
	destinoIni int
	origenIni  int
	cantidad   int
}

var seeds []int
var seedsC []int
var funcion [][]conversion

func main() {

	t = 20
	//t = 4
	n = 7

	inputs := []int{42, 49, 32, 47, 21, 37, 36}
	//inputs := []int{2, 3, 4, 2, 3, 2, 2}
	var basura string
	fmt.Scan(&basura)
	seeds = make([]int, t/2)
	seedsC = make([]int, t/2)
	for c := 0; c < t/2; c++ {
		fmt.Scan(&seeds[c])
		fmt.Scan(&seedsC[c])
	}

	funcion = make([][]conversion, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&basura)
		fmt.Scan(&basura)
		funcion[i] = make([]conversion, inputs[i])
		for j := 0; j < inputs[i]; j++ {
			var a, b, c int
			fmt.Scan(&a, &b, &c)
			//fmt.Println(a, b, c)
			funcion[i][j] = conversion{a, b, c}
		}
	}

	minimo := -1
	for c := 0; c < t/2; c++ {

		lim := seedsC[c]
		//fmt.Println("lim", lim)
		for k := 0; k < lim; k++ {
			//fmt.Println("seed", seeds[c]+k)
			seed := seeds[c] + k
			for i := 0; i < n; i++ {
				for j := 0; j < inputs[i]; j++ {
					//fmt.Println(funcion[i][j], seed)
					if funcion[i][j].origenIni <= seed && seed <= funcion[i][j].origenIni+funcion[i][j].cantidad-1 {

						seed = funcion[i][j].destinoIni + (seed - funcion[i][j].origenIni)
						break
					}
				}
			}

			if minimo == -1 || seed < minimo {

				minimo = seed
				//fmt.Println("minimo", seed)
			}
		}

	}
	fmt.Println(minimo)

}
