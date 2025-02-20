package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

var n int
var secretNumbers []int64
var secuences [][]int64
var secuencesDiff [][]int64
var pruneValue int64 = 16777216

func nextSecretNumber(number int64) int64 {
	number = ((number * 64) ^ number) % pruneValue
	number = ((number / 32) ^ number) % pruneValue
	number = ((number * 2048) ^ number) % pruneValue
	return number
}

// KMP busca todas las ocurrencias del patrón en el texto y retorna los índices donde empieza cada coincidencia.
func KMP(si int, pattern []int64, lps []int) int {

	i, j := 0, 0 // i para el texto, j para el patrón

	for i < len(secuencesDiff[si]) {
		if secuencesDiff[si][i] == pattern[j] {
			i++
			j++

			// Si encontramos el patrón completo
			if j == len(pattern) {
				return i
			}
		} else {
			if j > 0 {
				j = lps[j-1] // Reinicia j usando la tabla LPS
			} else {
				i++ // Avanza en el texto
			}
		}
	}

	return -1
}

// buildLPS construye la tabla de "longest prefix suffix" (LPS) para el patrón.
func buildLPS(sub []int64) []int {
	n := 4
	lps := make([]int, n)
	length := 0 // Longitud del prefijo más largo que es también sufijo
	i := 1

	for i < n {
		if sub[i] == sub[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length > 0 {
				length = lps[length-1] // Retrocede al valor previo en LPS
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}

func PrintSecuences(i, veces int) {
	for j := 0; j < veces+1; j++ {
		if j != 0 && secuencesDiff[i][j-1] < 0 {
			fmt.Printf(" %d ", secuences[i][j])
		} else {
			fmt.Printf("%d ", secuences[i][j])
		}
	}
	fmt.Println()
	fmt.Printf("  ")
	for j := 0; j < veces; j++ {
		fmt.Printf("%d ", secuencesDiff[i][j])
	}
	fmt.Println()
}

func main() {
	var total uint64 = 0
	tam := 4
	tam = 1749
	n = tam
	veces := 2000
	secretNumbers = make([]int64, tam)
	secuences = make([][]int64, tam)
	secuencesDiff = make([][]int64, tam)
	reader := bufio.NewReader(os.Stdin)
	for c := 0; c < tam; c++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		secretNumbers[c], _ = strconv.ParseInt(line, 10, 64)
	}

	for i := 0; i < n; i++ {
		original := secretNumbers[i]
		secuences[i] = make([]int64, veces+1)
		secuencesDiff[i] = make([]int64, veces)
		secuences[i][0] = original % 10
		for j := 1; j <= veces; j++ {
			secretNumbers[i] = nextSecretNumber(secretNumbers[i])
			secuences[i][j] = secretNumbers[i] % 10
			secuencesDiff[i][j-1] = secuences[i][j] - secuences[i][j-1]
		}
	}

	var sub []int64 = make([]int64, 4)
	var maximo int = 0

	//l,o,p,q
	maxPosible := 9
	for l := -9; l <= 9; l++ {
		for o := -9; o <= 9; o++ {
			if Abs(l+o) > 9 {
				continue
			}
			fmt.Println(l, o, maximo)
			for p := -9; p <= 9; p++ {
				if Abs(l+o+p) > 9 || Abs(o+p) > 9 {
					continue
				}
				for q := -9; q <= 9; q++ {
					if Abs(l+o+p+q) > 9 || Abs(o+p+q) > 9 || Abs(p+q) > 9 {
						continue
					}
					var nuevoValor int = 0
					sub[0] = int64(l)
					sub[1] = int64(o)
					sub[2] = int64(p)
					sub[3] = int64(q)
					lps := buildLPS(sub)
					x := l + o + p + q
					maxPosible = 9 + x
					//fmt.Println(l, o, p, q, maxPosible)
					for i := 0; i < n; i++ {
						if nuevoValor+(n-i)*min(maxPosible, 9) <= maximo {
							//fmt.Println("break")
							break
						}
						sum := KMP(i, sub, lps)
						if sum >= 0 {
							nuevoValor += int(secuences[i][sum])
						}
					}
					if nuevoValor > maximo {
						maximo = nuevoValor
					}
				}
			}
		}
	}

	total = uint64(maximo)
	fmt.Println(total)
}
