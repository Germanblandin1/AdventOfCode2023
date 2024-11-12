package main

import (
	"fmt"
	"sort"
)

var t int
var n, m int

type jugada struct {
	cartas2 string
	cartas  string
	tipo    int
	puntos  int
}

type PorTipoYCartas []jugada

func (a PorTipoYCartas) Len() int { return len(a) }

func (a PorTipoYCartas) Less(i, j int) bool {
	if a[i].tipo != a[j].tipo {
		return a[i].tipo < a[j].tipo
	}
	return a[i].cartas < a[j].cartas
}

func (a PorTipoYCartas) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func OrdenarPorTipoYCartas(jugadas []jugada) {
	sort.Sort(PorTipoYCartas(jugadas))
}

var jugadas [1001]jugada

var valores = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 1,
	'Q': 12,
	'K': 13,
	'A': 14,
}

// JJA62 5 291
func definitipo(cartas string) int {
	repetidas := make(map[rune]int)
	for i := 0; i < len(cartas); i++ {
		repetidas[rune(cartas[i])]++
	}

	max := 0
	index := -1
	for i := 0; i < len(cartas); i++ {
		if cartas[i] == 'J' {
			continue
		}
		if repetidas[rune(cartas[i])] > max {
			max = repetidas[rune(cartas[i])]
			index = i
		}
	}
	if index != -1 {
		cantidadJ := repetidas['J']
		repetidas[rune(cartas[index])] += cantidadJ

	}

	conteo := make(map[int]int)
	for letra, v := range repetidas {
		if letra == 'J' {
			continue
		}
		conteo[v]++
	}
	if conteo[5] == 1 || cartas == "JJJJJ" {
		return 7
	} else if conteo[4] == 1 {
		return 6
	} else if conteo[3] == 1 && conteo[2] == 1 {
		return 5
	} else if conteo[3] == 1 {
		return 4
	} else if conteo[2] == 2 {
		return 3
	} else if conteo[2] == 1 {
		return 2
	} else {
		return 1
	}

}

func main() {

	n = 1000

	for i := 0; i < n; i++ {
		fmt.Scan(&jugadas[i].cartas2)
		newcartas := make([]rune, 0)
		for j := 0; j < len(jugadas[i].cartas2); j++ {
			newcartas = append(newcartas, rune('A'+valores[rune(jugadas[i].cartas2[j])]))
		}
		fmt.Scan(&jugadas[i].puntos)
		jugadas[i].tipo = definitipo(jugadas[i].cartas2)
		jugadas[i].cartas = string(newcartas)

	}
	OrdenarPorTipoYCartas(jugadas[:n])
	total := 0
	for i := 0; i < n; i++ {
		fmt.Println(jugadas[i].cartas2, jugadas[i].tipo, jugadas[i].puntos)
		total += jugadas[i].puntos * (i + 1)
	}
	fmt.Println(total)

}
