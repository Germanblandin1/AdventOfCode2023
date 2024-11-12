package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var basura string
var t int
var n, m int
var INF int = 1000000000
var actual int

var matriz [][][]rune

func distancia(a, b string) int {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count
}

func main() {
	t = 2
	fmt.Scan(&t)
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	lim := t
	matriz = make([][][]rune, lim+1)
	c := 0
	matriz[0] = make([][]rune, 0)
	for i := 0; i < lim; i++ {
		basura := ""
		basura, _ = reader.ReadString('\n')
		basura = strings.Trim(basura, "\n")
		basura = strings.TrimSpace(basura)
		//fmt.Println(basura)
		//fmt.Println(len(basura))
		if len(basura) == 0 {
			c++
			matriz[c] = make([][]rune, 0)
		} else {
			matriz[c] = append(matriz[c], []rune(basura))
		}
	}
	c++

	suma := 0
	for t := 0; t < c; t++ {
		//busqueda vertical
		n = len(matriz[t])
		m = len(matriz[t][0])
		for j := 1; j < m; j++ {

			llevo := 0

			ant, sig := j-1, j
			esEspejo := true
			useComodin := false
			for esEspejo {
				anterior := ""
				actual := ""
				for i := 0; i < n; i++ {

					anterior += string(matriz[t][i][ant])
					actual += string(matriz[t][i][sig])
				}

				if !useComodin && distancia(anterior, actual) == 1 {
					useComodin = true
				} else if anterior == actual {
					llevo++
				} else {
					esEspejo = false
				}

				ant--
				sig++
				if ant < 0 || sig >= m {
					break
				}
			}
			if esEspejo && useComodin {
				//fmt.Println("vertifcak", t, j)
				suma += j
			}
		}

		for i := 1; i < n; i++ {

			llevo := 0

			ant, sig := i-1, i
			esEspejo := true
			useComodin := false
			for esEspejo {
				anterior := string(matriz[t][ant])
				actual := string(matriz[t][sig])

				if !useComodin && distancia(anterior, actual) == 1 {
					useComodin = true
				} else if anterior == actual {
					llevo++
				} else {
					esEspejo = false
				}
				ant--
				sig++
				if ant < 0 || sig >= n {
					break
				}
			}
			if esEspejo && useComodin {
				//fmt.Println("hori", t, i)
				suma += 100 * i
			}
		}
	}
	fmt.Println(suma)

}
