package main

import "fmt"

var basura string
var t int
var n, m int
var INF int = 1000000000
var actual int

var matriz [][]rune

// baja, derecha, arriba, izquierda
var movi = []int{1, 0, -1, 0}
var movj = []int{0, 1, 0, -1}

func simulate(i_ini, j_ini, dir int) {
	i := i_ini
	j := j_ini
	ori_i := i
	ori_j := j
	for {
		if i < 0 || i >= n || j < 0 || j >= m {
			break
		}
		if matriz[i][j] == 'O' {
			matriz[i][j] = '.'
			matriz[ori_i][ori_j] = 'O'
			ori_i = ori_i + movi[dir]
			ori_j = ori_j + movj[dir]
		}
		if matriz[i][j] == '#' {
			ori_i = i + movi[dir]
			ori_j = j + movj[dir]
		}
		i += movi[dir]
		j += movj[dir]

	}
}

func printMatriz() {
	for i := 0; i < n; i++ {
		fmt.Println(string(matriz[i]))
	}
	fmt.Println()

}
func calcularMatriz() int {
	suma := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matriz[i][j] == 'O' {
				suma += n - i
			}
		}
	}
	return suma
}

func cycle() {
	// orden: norte, oeste,sur,este

	//inclina hacia el norte
	for j := 0; j < m; j++ {
		simulate(0, j, 0)
	}
	//fmt.Println("despues de norte")
	//printMatriz()
	//inclina hacia el oeste
	for i := 0; i < n; i++ {
		simulate(i, 0, 1)
	}
	//fmt.Println("despues de oeste")
	//printMatriz()
	//inclina hacia el sur
	for j := 0; j < m; j++ {
		simulate(n-1, j, 2)
	}
	//fmt.Println("despues de sur")
	//printMatriz()
	//inclina hacia el este
	for i := 0; i < n; i++ {
		simulate(i, m-1, 3)
	}
	//fmt.Println("despues de este")
	//printMatriz()

}

func main() {
	fmt.Scan(&n, &m)
	matriz = make([][]rune, n)
	for i := 0; i < n; i++ {
		matriz[i] = make([]rune, m)
		fmt.Scan(&basura)
		matriz[i] = []rune(basura)
	}
	printMatriz()
	numcicles := 1000
	for i := 0; i < numcicles; i++ {
		//fmt.Println("ciclo", i, "de", numcicles, "ciclos")
		cycle()
		fmt.Println("Calculo de matriz en ciclo", i+1, ":", calcularMatriz())
		//fmt.Println()
		//fmt.Println()
	}

	suma := 0
	// for j := 0; j < m; j++ {
	// 	desde := n
	// 	newdesde := n
	// 	llevo := 0
	// 	for i := 0; i < n; i++ {
	// 		if matriz[i][j] == 'O' {
	// 			llevo++
	// 		}
	// 		if matriz[i][j] == '#' || i == n-1 {
	// 			lim := desde
	// 			lim2 := desde - llevo
	// 			//fmt.Println("j", j, "i", i, "lim", lim, "lim2", lim2)
	// 			//fmt.Println(((lim * (lim + 1)) / 2), ((lim2 * (lim2 + 1)) / 2))
	// 			suma += ((lim * (lim + 1)) / 2) - ((lim2 * (lim2 + 1)) / 2)
	// 			desde = newdesde - 1
	// 			llevo = 0
	// 		}
	// 		newdesde--
	// 	}
	// }
	fmt.Println(suma)

}
