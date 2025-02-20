package main

import "fmt"

var t int
var n, m int

func main() {

	t = 100
	//t = 4
	var basura string

	total := 0
	var valor int
	for c := 0; c < t; c++ {
		fmt.Scan(&basura)
		fmt.Scan(&valor)
		var numero int
		var color string

		totalR := 0
		totalB := 0
		totalG := 0
		maximoR := 0
		maximoB := 0
		maximoG := 0
		for {
			fmt.Scan(&numero)
			fmt.Scan(&color)
			//fmt.Println(numero, color)
			if color[0] == 'r' {
				totalR += numero
			} else if color[0] == 'b' {
				totalB += numero
			} else if color[0] == 'g' {
				totalG += numero
			}

			if color[len(color)-1] == ';' || (color[len(color)-1] != ',' && color[len(color)-2] != ';') {
				if totalR > maximoR {
					maximoR = totalR
				}
				if totalB > maximoB {
					maximoB = totalB
				}
				if totalG > maximoG {
					maximoG = totalG
				}
				totalB = 0
				totalR = 0
				totalG = 0
			}

			if color[len(color)-1] != ';' && color[len(color)-1] != ',' {
				break
			}
		}
		total += (maximoR * maximoB * maximoG)
	}
	fmt.Println(total)

}
