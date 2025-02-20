package main

import (
	"fmt"
	"strconv"
	"strings"
)

var basura string
var t int
var n, m int
var INF int = 1000000000
var actual int

type Pair struct {
	x, y int
}

var pares []Pair

func area_gauss(pares []Pair) int {
	var area int
	//pares = append(pares, Pair{pares[0].x, pares[0].y})
	A := 0
	for i := 0; i < len(pares)-1; i++ {
		A += pares[i].x * pares[i+1].y
	}
	B := 0
	for i := 0; i < len(pares)-1; i++ {
		B += pares[i+1].x * pares[i].y
	}
	fmt.Println(A, B)
	if B > A {
		B, A = A, B
	}
	area = (A - B) / 2
	return area
}

func main() {
	fmt.Scan(&t)

	pos_i := 0
	pos_j := 0

	pares = make([]Pair, 0)
	pares = append(pares, Pair{0, 0})
	perimetro := 0
	for c := 0; c < t; c++ {
		var strDir string
		var pasos int
		var color string

		fmt.Scan(&strDir, &pasos, &color)
		color = strings.Trim(color, "()")
		color = strings.Trim(color, "#")
		strNumero := color[:len(color)-1]
		newpasos, _ := strconv.ParseInt(strNumero, 16, 64)
		pasos = int(newpasos)
		strDir = color[len(color)-1:]
		//fmt.Println(strDir, pasos)
		//pasos++
		// R: 0, D: 1, L: 2, U: 3
		if strDir == "0" {
			pos_i += pasos
		} else if strDir == "2" {
			pos_i -= pasos
		} else if strDir == "1" {
			pos_j -= pasos
		} else if strDir == "3" {
			pos_j += pasos
		}
		perimetro += pasos
		pares = append(pares, Pair{pos_i, pos_j})
	}

	fmt.Println(area_gauss(pares) + perimetro/2 + 1)
}
