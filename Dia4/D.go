package main

import (
	"fmt"
)

var t int
var n, m int

var ganadoras map[int]bool

var instancias map[int]uint64

func main() {

	t = 201
	n = 10
	m = 25
	var total uint64 = 0
	instancias = make(map[int]uint64)
	for c := 0; c < t; c++ {
		instancias[c]++
		ganadoras = make(map[int]bool)
		var car string

		fmt.Scan(&car)
		fmt.Scan(&car)

		var a int
		for i := 0; i < n; i++ {
			fmt.Scan(&a)
			ganadoras[a] = true
		}
		fmt.Scan(&car)

		count := 0
		for i := 0; i < m; i++ {
			fmt.Scan(&a)
			_, ok := ganadoras[a]
			if ok {
				count++
			}
		}
		if count > 0 {
			for k := 1; k <= count; k++ {
				instancias[c+k] += instancias[c]
			}
		}
	}
	for i := 0; i < t; i++ {
		total += instancias[i]
	}
	fmt.Println(total)

}
