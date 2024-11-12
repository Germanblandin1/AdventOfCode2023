package main

import (
	"fmt"
	"strings"
)

var basura string
var t int
var n, m int
var mapa map[string]int
var mapaINV map[int]string
var grafo [][]int

var origenes []int
var destinos []int

var marca []bool

func simulate(moves string) int {

	i := 0
	count := 0
	llego := 0
	fmt.Println(origenes, destinos)
	for {
		fmt.Println(origenes, count)
		if llego == len(destinos) {
			return count
		}

		for j := 0; j < len(origenes); j++ {
			origenes[j] = grafo[origenes[j]][moves[i]-'0']
		}
		marca := make([]bool, n)
		llego = 0
		for j := 0; j < len(origenes); j++ {
			for k := 0; k < len(destinos); k++ {
				if origenes[j] == destinos[k] && !marca[destinos[k]] {
					marca[destinos[k]] = true
					llego++
				}
			}
		}

		count++
		i = (i + 1) % len(moves)
	}

}

var acumuladas [][][]uint64
var ciclos []uint64
var totales []uint64
var primeros []uint64
var indices []int
var acumuladadestinos []uint64

func simulateSimple(moves string, origen, destino int) (uint64, uint64, int, uint64) {

	i := 0
	count := uint64(0)
	//fmt.Println(origen, destino)
	pos := origen
	fmt.Println(origen, destino)
	acumuladadestino := uint64(0)
	for {
		if pos == destino {
			fmt.Println("llego", mapaINV[pos], mapaINV[destino], count, acumuladas[origen][i][pos], i)
			acumuladadestino = count
		}
		if acumuladas[origen][i][pos] != 0 {
			fmt.Println("ACUMMMd", mapaINV[pos], mapaINV[destino], count, acumuladas[origen][i][pos], i)
			return count, acumuladas[origen][i][pos], i, acumuladadestino
		}
		acumuladas[origen][i][pos] = count
		pos = grafo[pos][moves[i]-'0']
		count++
		i = (i + 1) % len(moves)
	}

}

func main() {

	n = 3
	//n = 7
	//n = 8
	//m = 2
	n = 730
	m = 6
	moves := ""
	fmt.Scan(&moves)
	moves = strings.ReplaceAll(moves, "L", "0")
	moves = strings.ReplaceAll(moves, "R", "1")
	//fmt.Println(moves)
	mapa = make(map[string]int)
	mapaINV = make(map[int]string)

	grafo = make([][]int, n)
	origenes = make([]int, m)
	destinos = make([]int, m)
	for i := 0; i < n; i++ {
		grafo[i] = make([]int, 2)
	}
	index := 0
	for i := 0; i < n; i++ {
		origen := ""
		fmt.Scan(&origen)
		fmt.Scan(&basura)
		destino1 := ""
		destino2 := ""
		fmt.Scan(&destino1)
		fmt.Scan(&destino2)

		destino1 = strings.TrimPrefix(destino1, "(")
		destino1 = strings.TrimSuffix(destino1, ",")
		destino2 = strings.TrimSuffix(destino2, ")")
		//fmt.Println(origen, destino1, destino2)

		if _, ok := mapa[origen]; !ok {
			mapa[origen] = index
			mapaINV[index] = origen
			index++

		}
		if _, ok := mapa[destino1]; !ok {
			mapa[destino1] = index
			mapaINV[index] = destino1
			index++
		}
		if _, ok := mapa[destino2]; !ok {
			mapa[destino2] = index
			mapaINV[index] = destino2
			index++
		}
		grafo[mapa[origen]][0] = mapa[destino1]
		grafo[mapa[origen]][1] = mapa[destino2]

	}
	origenes[0] = mapa["XSA"]
	origenes[1] = mapa["VVA"]
	origenes[2] = mapa["TTA"]
	origenes[3] = mapa["AAA"]
	origenes[4] = mapa["NBA"]
	origenes[5] = mapa["MHA"]
	destinos[0] = mapa["TKZ"]
	destinos[1] = mapa["PSZ"]
	destinos[2] = mapa["RFZ"]
	destinos[3] = mapa["ZZZ"]
	destinos[4] = mapa["HGZ"]
	destinos[5] = mapa["GJZ"]

	// origenes[0] = mapa["11A"]
	// origenes[1] = mapa["22A"]
	// destinos[0] = mapa["11Z"]
	// destinos[1] = mapa["22Z"]

	ciclos = make([]uint64, n)
	acumuladas = make([][][]uint64, n)
	for i := 0; i < n; i++ {
		acumuladas[i] = make([][]uint64, len(moves))
		for j := 0; j < len(moves); j++ {
			acumuladas[i][j] = make([]uint64, n)
		}
	}
	totales = make([]uint64, n)
	primeros = make([]uint64, n)
	indices = make([]int, n)
	acumuladadestinos = make([]uint64, n)
	maximo := uint64(0)
	for i := 0; i < len(origenes); i++ {
		total, primersitio, indexori, acumuladadestino := simulateSimple(moves, origenes[i], destinos[i])
		if primersitio > maximo {
			maximo = uint64(primersitio)
		}
		ciclos[origenes[i]] = uint64(total) - uint64(primersitio)
		totales[origenes[i]] = uint64(total)
		primeros[origenes[i]] = uint64(primersitio)
		indices[origenes[i]] = indexori
		acumuladadestinos[origenes[i]] = acumuladadestino
		//fmt.Println(acumuladas[origenes[i]][destinos[i]])
	}

	for i := 0; i < len(origenes); i++ {
		fmt.Println("total", totales[origenes[i]], "acumulada", acumuladadestinos[origenes[i]], "maxmo", maximo, "primer", primeros[origenes[i]], "ciclo", ciclos[origenes[i]])
		fmt.Println("asub", ((totales[origenes[i]]-acumuladadestinos[origenes[i]])+(maximo-primeros[origenes[i]]))%ciclos[origenes[i]], "nsub", ciclos[origenes[i]])
	}
	fmt.Println(maximo)

	//fmt.Println(simulate(moves))
	//XSA TKZ 16409
	//VVA PSZ 12643
	//TTA RFZ 21251
	//AAA ZZZ 15871
	//NBA HGZ 19637
	//MHA GJZ 11567

}
