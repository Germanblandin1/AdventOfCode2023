package main

import (
	"fmt"
	"sort"
)

var basura string
var t int
var n, m int
var INF int = 1000000000

type point struct {
	x, y, z int
}

type line struct {
	p1, p2 point
}

var lines []line
var linesCopia []line
var linesOriginal []line

func ordenarLines(lines []line) {
	sort.Slice(lines, func(i, j int) bool {

		if lines[i].p1.z != lines[j].p1.z {
			return lines[i].p1.z < lines[j].p1.z
		}

		if lines[i].p1.x != lines[j].p1.x {
			return lines[i].p1.x < lines[j].p1.x
		}

		return lines[i].p1.y < lines[j].p1.y
	})
}

const MAXX = 11
const MAXY = 11
const MAXZ = 400

var box [MAXX][MAXY][MAXZ]bool
var boxCopia [MAXX][MAXY][MAXZ]bool
var boxOriginal [MAXX][MAXY][MAXZ]bool

func estaBloqueado(linea line, pisoZ int) bool {
	for i := linea.p1.x; i <= linea.p2.x; i++ {
		for j := linea.p1.y; j <= linea.p2.y; j++ {
			if box[i][j][pisoZ] == true {
				return true
			}
		}
	}
	return false
}

func LimpiarBox() {
	for i := 0; i < MAXX; i++ {
		for j := 0; j < MAXY; j++ {
			for k := 0; k < MAXZ; k++ {
				if k == 0 {
					continue
				}
				box[i][j][k] = false
			}
		}
	}
}

func marcarBloque(linea line, pisoZ int) {
	tamZ := linea.p2.z - linea.p1.z + 1
	// una vez caido el bloque marcarlo en la caja
	for i := linea.p1.x; i <= linea.p2.x; i++ {
		for j := linea.p1.y; j <= linea.p2.y; j++ {
			for k := pisoZ; k <= pisoZ+tamZ-1; k++ {
				box[i][j][k] = !box[i][j][k]
			}
		}
	}
}

func simularUnBloque(linea *line, cambiar bool) {
	//hacer que caiga el bloque en z hacia 0
	pisoZ := linea.p1.z
	for {
		if estaBloqueado(*linea, pisoZ-1) {
			break
		}
		pisoZ--
	}

	marcarBloque(*linea, pisoZ)
	if cambiar {
		tamZ := linea.p2.z - linea.p1.z + 1
		linea.p1.z = pisoZ
		linea.p2.z = pisoZ + tamZ - 1
	}

}

func simularCaida(fuera int, cambiar bool) {
	for i := 0; i < n; i++ {
		if i == fuera {
			continue
		}
		simularUnBloque(&lines[i], cambiar)
	}
}

func copiarBox(original [MAXX][MAXY][MAXZ]bool, copia *[MAXX][MAXY][MAXZ]bool) {
	for i := 0; i < MAXX; i++ {
		for j := 0; j < MAXY; j++ {
			for k := 0; k < MAXZ; k++ {
				copia[i][j][k] = original[i][j][k]
			}
		}
	}
}

func compararBox(original [MAXX][MAXY][MAXZ]bool, copia [MAXX][MAXY][MAXZ]bool) bool {
	for i := 0; i < MAXX; i++ {
		for j := 0; j < MAXY; j++ {
			for k := 0; k < MAXZ; k++ {
				if original[i][j][k] != copia[i][j][k] {
					return false
				}
			}
		}
	}
	return true

}

func mostrarBox(box [MAXX][MAXY][MAXZ]bool) {
	for k := 0; k < 8; k++ {
		for i := 0; i < MAXX; i++ {
			for j := 0; j < MAXY; j++ {
				if box[i][j][k] {
					fmt.Print("1")
				} else {
					fmt.Print("0")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func copiarLines(original []line, copia []line) {
	for i := 0; i < n; i++ {
		copia[i].p1.x = original[i].p1.x
		copia[i].p1.y = original[i].p1.y
		copia[i].p1.z = original[i].p1.z
		copia[i].p2.x = original[i].p2.x
		copia[i].p2.y = original[i].p2.y
		copia[i].p2.z = original[i].p2.z
	}
}

func main() {
	fmt.Scan(&n)

	lines = make([]line, n)
	linesCopia = make([]line, n)
	linesOriginal = make([]line, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&lines[i].p1.x, &lines[i].p1.y, &lines[i].p1.z, &lines[i].p2.x, &lines[i].p2.y, &lines[i].p2.z)
	}

	for i := 0; i < MAXX; i++ {
		for j := 0; j < MAXY; j++ {
			box[i][j][0] = true
		}
	}

	ordenarLines(lines)
	for i := 0; i < n; i++ {
		fmt.Println(lines[i])
	}
	fmt.Println("despues")
	simularCaida(-1, true)
	for i := 0; i < n; i++ {
		fmt.Println(lines[i])
	}
	copiarBox(box, &boxOriginal)
	ordenarLines(lines)
	copiarLines(lines, linesOriginal)
	//mostrarBox(box)

	suma := 0
	acum := 0
	for i := 0; i < n; i++ {
		marcarBloque(lines[i], lines[i].p1.z)
		copiarBox(box, &boxCopia)
		ordenarLines(lines)
		//fmt.Println("Quitando", lines[i])
		fmt.Println("Antes:")
		//mostrarBox(box)
		LimpiarBox()

		simularCaida(i, true)
		fmt.Println("quedo asi:")
		//mostrarBox(box)
		if compararBox(box, boxCopia) {
			fmt.Println("quedaron iguales")
			suma++
		} else {
			fmt.Println("quedaron diferentes")
			for j := 0; j < n; j++ {
				if j == i {
					continue
				}
				//si las lineas son diferentes acumular
				if lines[j].p1.x != linesOriginal[j].p1.x || lines[j].p1.y != linesOriginal[j].p1.y || lines[j].p1.z != linesOriginal[j].p1.z || lines[j].p2.x != linesOriginal[j].p2.x || lines[j].p2.y != linesOriginal[j].p2.y || lines[j].p2.z != linesOriginal[j].p2.z {
					acum++
				}
			}
		}
		copiarBox(boxOriginal, &box)
		copiarLines(linesOriginal, lines)
	}
	fmt.Println(suma, acum)

}
