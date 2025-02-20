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
var inicio int

type Rule struct {
	determinante int
	signo        string
	valor        int
	produccion   int
}

var mapa map[string]int

var reglas [][]Rule

// x=0,m=1,a=2,s=3
var valores [][]int

func mapear(s string) int {
	if val, ok := mapa[s]; ok {
		return val
	}
	mapa[s] = actual
	actual++
	return actual - 1
}

func mapearValor(s string) int {
	// x,m,a,s
	if s == "x" {
		return 0
	}
	if s == "m" {
		return 1
	}
	if s == "a" {
		return 2
	}
	if s == "s" {
		return 3
	}
	return -1
}

func parseStringRule(s string) {
	//dividir primero por {
	primerCorte := strings.Split(s, "{")
	//fmt.Println(primerCorte, len(primerCorte))
	origenString := primerCorte[0]
	origen := mapear(origenString)
	resto := strings.Trim(primerCorte[1], "}")
	segundoCorte := strings.Split(resto, ",")
	if origenString == "in" {
		inicio = origen
	}
	reglas[origen] = make([]Rule, 0)
	for i := 0; i < len(segundoCorte); i++ {

		tercerCorte := strings.Split(segundoCorte[i], ":")
		if len(tercerCorte) == 2 {
			condicion := tercerCorte[0]
			produccionStr := tercerCorte[1]
			produccion := mapear(produccionStr)
			simbolo := "<"
			if strings.Contains(condicion, ">") {
				simbolo = ">"
			}
			cuartoCorte := strings.Split(condicion, simbolo)
			determinanteStr := cuartoCorte[0]
			determinante := mapearValor(determinanteStr)
			valorStr := cuartoCorte[1]
			valor, _ := strconv.Atoi(valorStr)
			reglas[origen] = append(reglas[origen], Rule{determinante, simbolo, valor, produccion})

		} else {
			condicionFinal := tercerCorte[0]
			origenFinal := mapear(condicionFinal)
			reglas[origen] = append(reglas[origen], Rule{-1, "final", 0, origenFinal})
		}
	}

}

func parseStringValor(s string, i int) {
	s = strings.Trim(s, "{}")
	pedazo := strings.Split(s, ",")
	for j := 0; j < 4; j++ {
		cortar := strings.Split(pedazo[j], "=")
		valor, _ := strconv.Atoi(cortar[1])
		valores[i][mapearValor(cortar[0])] = valor
	}

}

func simular(t int) bool {
	nodoActual := inicio
	encontro := false
	aceptado := false
	for !encontro {

		for i := 0; i < len(reglas[nodoActual]); i++ {
			rule := reglas[nodoActual][i]
			if rule.signo == "final" {
				nodoActual = rule.produccion
				break
			}
			esmenor := valores[t][rule.determinante] < rule.valor
			esmayor := valores[t][rule.determinante] > rule.valor
			var condicion bool = esmayor
			if rule.signo == "<" {
				condicion = esmenor
			}
			if condicion {
				nodoActual = rule.produccion
				break
			}
		}
		// R es 0, A es 1
		if nodoActual == 0 {
			encontro = true
		}
		if nodoActual == 1 {
			encontro = true
			aceptado = true
		}

	}
	return aceptado

}

func sumarvalores(i int) int {
	return valores[i][0] + valores[i][1] + valores[i][2] + valores[i][3]
}

//	type Rule struct {
//		determinante int
//		signo        string
//		valor        int
//		produccion   int
//	}
func calcularBeta(nodo int, rangos [][]int) uint64 {

	if nodo == 0 {
		return 0
	}

	for i := 0; i < 4; i++ {
		if rangos[i][0] > rangos[i][1] {
			return 0
		}
	}

	//llegue a A
	fmt.Println(rangos)
	var suma uint64 = 0
	if nodo == 1 {
		fmt.Println("llegue")
		suma = 1
		for i := 0; i < 4; i++ {
			valor := uint64(rangos[i][1] - rangos[i][0] + 1)
			suma *= valor
		}
		fmt.Println(suma)
		return suma
	}

	respaldo := make([][]int, 4)
	for i := 0; i < 4; i++ {
		respaldo[i] = make([]int, 2)
	}

	for i := 0; i < len(reglas[nodo]); i++ {
		for i := 0; i < 4; i++ {
			respaldo[i][0] = rangos[i][0]
			respaldo[i][1] = rangos[i][1]
		}
		rule := reglas[nodo][i]
		if rule.signo == ">" {
			rangos[rule.determinante][0] = rule.valor + 1
		} else if rule.signo == "<" {
			rangos[rule.determinante][1] = rule.valor - 1
		}

		beta := calcularBeta(rule.produccion, rangos)

		if beta != 0 {
			suma += beta
		}

		for i := 0; i < 4; i++ {
			rangos[i][0] = respaldo[i][0]
			rangos[i][1] = respaldo[i][1]
		}
		if rule.signo == ">" {
			//rangos[rule.determinante][0] = rule.valor + 1
			rangos[rule.determinante][1] = rule.valor
		} else if rule.signo == "<" {
			//rangos[rule.determinante][1] = rule.valor - 1
			rangos[rule.determinante][0] = rule.valor
		}

	}

	return suma

}

func main() {
	fmt.Scan(&n, &m)
	mapa = make(map[string]int)
	reglas = make([][]Rule, n+2)
	valores = make([][]int, m)
	for i := 0; i < m; i++ {
		valores[i] = make([]int, 4)
	}
	actual = 0
	mapear("R")
	mapear("A")

	for i := 0; i < n; i++ {
		fmt.Scan(&basura)
		parseStringRule(basura)
	}

	for i := 0; i < m; i++ {
		fmt.Scan(&basura)
		parseStringValor(basura, i)
	}

	rangos := make([][]int, 4)
	for i := 0; i < 4; i++ {
		rangos[i] = make([]int, 2)
		rangos[i][0] = 1
		rangos[i][1] = 4000
	}

	fmt.Println(calcularBeta(inicio, rangos))

}
