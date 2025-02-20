package main

import "fmt"

var basura string
var t int
var n, m int
var INF int = 1000000000

var pasos int

var movi = [4]int{1, 0, -1, 0}
var movj = [4]int{0, 1, 0, -1}
var ini_i, ini_j int

var matriz [][]rune
var matriz_exp [][]rune

type Pair struct {
	i, j, pasos int
}

type Queue []Pair

func (q *Queue) Enqueue(p Pair) {
	*q = append(*q, p)
}

func (q *Queue) Dequeue() (Pair, bool) {
	if len(*q) == 0 {
		return Pair{}, false
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, true
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

var dist [][]int
var marca [][]bool

func bfs(ini_i, ini_j int, pasos int) int {

	dist = make([][]int, n)
	marca = make([][]bool, n)

	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		marca[i] = make([]bool, m)
		for j := 0; j < m; j++ {
			dist[i][j] = INF
		}
	}

	cola := Queue{}
	cola.Enqueue(Pair{ini_i, ini_j, 0})
	dist[ini_i][ini_j] = 0
	cuenta := 0
	for !cola.IsEmpty() {
		p, _ := cola.Dequeue()
		i := p.i
		j := p.j
		//fmt.Println(i, j, p.pasos)

		if p.pasos%2 == pasos%2 && marca[i][j] == false {
			cuenta++
			marca[i][j] = true
		}

		for k := 0; k < 4; k++ {
			ii := i + movi[k]
			jj := j + movj[k]
			if ii >= 0 && ii < n && jj >= 0 && jj < m && matriz[ii][jj] != '#' {
				if p.pasos+1 > pasos {
					continue
				}
				if dist[i][j]+1 < dist[ii][jj] {
					cola.Enqueue(Pair{ii, jj, p.pasos + 1})
					dist[ii][jj] = dist[i][j] + 1
				}

			}
		}
	}
	return cuenta
}

func main() {

	fmt.Scan(&n, &m)
	fmt.Scan(&ini_i, &ini_j)
	fmt.Scan(&pasos)
	matriz = make([][]rune, n)
	//dist = make([][]int, n)
	//marca = make([][]bool, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&basura)
		matriz[i] = []rune(basura)
		//dist[i] = make([]int, m)
		//marca[i] = make([]bool, m)
	}

	k := 39
	n_exp := n * k
	m_exp := m * k
	matriz_exp = make([][]rune, n_exp)
	for i := 0; i < n_exp; i++ {
		matriz_exp[i] = make([]rune, m_exp)
		for j := 0; j < m_exp; j++ {
			matriz_exp[i][j] = matriz[i%n][j%m]
		}
	}
	matriz = matriz_exp
	n = n_exp
	m = m_exp
	limite := 911 + 132 + 132
	ini_i = ini_i*k + k/2
	ini_j = ini_j*k + k/2

	// for i := 0; i < n; i++ {
	// 	fmt.Println(string(matriz[i]))
	// }

	valores := make([]int, limite+1)
	diff := make([]uint64, limite+1)

	ant := 0
	for i := 0; i < limite+1; i++ {
		cuenta := bfs(ini_i, ini_j, i)
		fmt.Println("pasos:", i, cuenta, "dif:", cuenta-ant)
		valores[i] = cuenta
		diff[i] = uint64(cuenta - ant)
		ant = cuenta
	}

	empiezaciclo := 911
	tamanociclo := 131
	valorAcalcular := 26501365

	fmt.Println(valores[empiezaciclo], diff[empiezaciclo], valores[empiezaciclo+tamanociclo], diff[empiezaciclo+tamanociclo])
	baseDiff := make([]uint64, tamanociclo+2)
	for i := 0; i < tamanociclo; i++ {
		baseDiff[i] = diff[empiezaciclo+i+tamanociclo] - diff[empiezaciclo+i]
		fmt.Println(baseDiff[i])
	}
	valorBase := uint64(valores[empiezaciclo]) - diff[empiezaciclo]
	fmt.Println(valores[empiezaciclo], baseDiff[0], valorBase)
	valorAcalcular -= empiezaciclo
	veces := uint64(valorAcalcular / tamanociclo)
	resto := valorAcalcular % tamanociclo
	valorAcalcular -= resto

	var sumaDiff uint64 = 0
	for i := 0; i < tamanociclo; i++ {
		sumaDiff += diff[empiezaciclo+i]*veces + baseDiff[i]*(veces*(veces-1)/2)
	}

	//sumaDiff += uint64(valorBase) + diff[empiezaciclo] + baseDiff[0]*veces
	sumaDiff += uint64(valorBase)
	for i := 0; i < resto+1; i++ {
		sumaDiff += diff[empiezaciclo+i] + baseDiff[i]*veces
		fmt.Println("suma:", sumaDiff)
	}

}
