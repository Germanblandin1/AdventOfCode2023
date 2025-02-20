package main

import (
	"fmt"
	"math"
)

var basura string
var t int
var n, m int

var matriz [][]rune

// -L|F7
// 7S-7|
// L|7||
// -L-J|
// L|-JF

var caracteres = []rune{'-', '|', 'L', 'J', 'F', '7', 'S'}
var movi = []int{-1, 0, 1, 0}
var movj = []int{0, 1, 0, -1}

// arriba:0 derecha:1 abajo:2 izquierda:3

var entradas map[rune][2]int = map[rune][2]int{
	'-': {1, 3},
	'|': {0, 2},
	'L': {0, 1},
	'J': {0, 3},
	'F': {2, 1},
	'7': {2, 3},
}

var distancias [][]int
var marca [][]int

type Pair struct {
	i, j, dir int
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

func bfs(i, j int) {

	distancias[i][j] = 0
	//fmt.Println(i, j, string(matriz[i][j]))

	cola := Queue{}
	cola.Enqueue(Pair{i, j, entradas[matriz[i][j]][0]})
	cola.Enqueue(Pair{i, j, entradas[matriz[i][j]][1]})

	for !cola.IsEmpty() {
		pair, _ := cola.Dequeue()
		i := pair.i
		j := pair.j
		dir := pair.dir
		//fmt.Println(i, j, string(matriz[i][j]), dir)

		ni := i + movi[dir]
		nj := j + movj[dir]
		nextCell := matriz[ni][nj]
		contra := (dir + 2) % 4
		nextDir := entradas[nextCell][0]
		if nextDir == contra {
			nextDir = entradas[nextCell][1]
		}
		//fmt.Println("---", ni, nj, string(nextCell), nextDir)
		if distancias[i][j]+1 < distancias[ni][nj] {
			distancias[ni][nj] = distancias[i][j] + 1
			cola.Enqueue(Pair{ni, nj, nextDir})
		}
	}

}

func dfs(i, j, index int) int {
	marca[i][j] = index
	can := 1
	for k := 0; k < 4; k++ {
		ni := i + movi[k]
		nj := j + movj[k]
		if ni >= 0 && ni < n && nj >= 0 && nj < m {
			if marca[ni][nj] == -1 && distancias[ni][nj] == 1000000000 {
				can += dfs(ni, nj, index)
			}
		}
	}
	return can
}

type Point struct {
	X, Y float64
}

type Line struct {
	P1, P2 Point
}

var Poligono []Point

func dfs2(i, j, index, dir int) int {
	marca[i][j] = index
	if matriz[i][j] != '-' && matriz[i][j] != '|' {
		Poligono = append(Poligono, Point{float64(j), float64(i)})
	}

	can := 1

	ni := i + movi[dir]
	nj := j + movj[dir]
	nextCell := matriz[ni][nj]
	contra := (dir + 2) % 4
	nextDir := entradas[nextCell][0]
	if nextDir == contra {
		nextDir = entradas[nextCell][1]
	}
	if marca[ni][nj] == -1 && distancias[ni][nj] != 1000000000 {
		can += dfs2(ni, nj, index, nextDir)
	}

	return can
}

func estaDentro(oi, oj, dir int) int {
	i := oi
	j := oj
	can := 0
	for {
		if !(i >= 0 && i < n && j >= 0 && j < m) {
			break
		}
		if distancias[i][j] != 1000000000 {
			can++
		}
		i += 1
		j += 1
	}
	return can % 2
}

func onLine(l Line, p Point) bool {
	// Check whether p is on the line or not
	if p.X <= math.Max(float64(l.P1.X), float64(l.P2.X)) &&
		p.X >= math.Min(float64(l.P1.X), float64(l.P2.X)) &&
		p.Y <= math.Max(float64(l.P1.Y), float64(l.P2.Y)) &&
		p.Y >= math.Min(float64(l.P1.Y), float64(l.P2.Y)) {
		return true
	}

	return false
}

func direction(a, b, c Point) int {
	val := (b.Y-a.Y)*(c.X-b.X) - (b.X-a.X)*(c.Y-b.Y)

	if val == 0 {
		fmt.Println((b.Y-a.Y)*(c.X-b.X), (b.X-a.X)*(c.Y-b.Y))
		fmt.Println(a, b, c, val)
		// Collinear
		return 0
	} else if val < 0 {
		// Anti-clockwise direction
		return 2
	}
	// Clockwise direction
	return 1
}

func isIntersect(l1, l2 Line) bool {
	// Four direction for two lines and points of other line
	dir1 := direction(l1.P1, l1.P2, l2.P1)
	dir2 := direction(l1.P1, l1.P2, l2.P2)
	dir3 := direction(l2.P1, l2.P2, l1.P1)
	dir4 := direction(l2.P1, l2.P2, l1.P2)

	// When intersecting
	if dir1 != dir2 && dir3 != dir4 {
		return true
	}

	// When p2 of line2 are on the line1
	if dir1 == 0 && onLine(l1, l2.P1) {
		return true
	}

	// When p1 of line2 are on the line1
	if dir2 == 0 && onLine(l1, l2.P2) {
		return true
	}

	// When p2 of line1 are on the line2
	if dir3 == 0 && onLine(l2, l1.P1) {
		return true
	}

	// When p1 of line1 are on the line2
	if dir4 == 0 && onLine(l2, l1.P2) {
		return true
	}

	return false
}

func checkInside(poly []Point, n int, p Point) bool {
	// When polygon has less than 3 edge, it is not polygon
	if n < 3 {
		return false
	}

	// Create a point at infinity, y is same as point p
	exline := Line{p, Point{9999, p.Y}}
	count := 0
	i := 0
	for {
		// Forming a line from two consecutive points of poly
		side := Line{poly[i], poly[(i+1)%n]}
		if isIntersect(side, exline) {
			// If side is intersects exline
			if direction(side.P1, p, side.P2) == 0 {
				fmt.Println(side.P1, p, side.P2)
				fmt.Println("collinear")
				return onLine(side, p)
			}
			count++
		}
		i = (i + 1) % n
		if i == 0 {
			break
		}
	}

	// When count is odd
	return count%2 == 1
}

func main() {

	fmt.Scan(&n)
	fmt.Scan(&m)
	matriz = make([][]rune, n)
	si, sj := -1, -1

	fmt.Scan(&si, &sj)
	distancias = make([][]int, n)
	marca = make([][]int, n)
	for i := 0; i < n; i++ {
		marca[i] = make([]int, m)
		distancias[i] = make([]int, m)
		fmt.Scan(&basura)
		matriz[i] = []rune(basura)
		for j := 0; j < m; j++ {
			marca[i][j] = -1
			distancias[i][j] = 1000000000
		}
	}

	maximo := 0
	bfs(si, sj)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if distancias[i][j] != 1000000000 && distancias[i][j] > maximo {
				maximo = distancias[i][j]
			}
		}
	}

	var comPos []Pair = make([]Pair, 0)
	var comCan []int = make([]int, 0)

	canCompo := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if distancias[i][j] == 1000000000 && marca[i][j] == -1 {
				comPos = append(comPos, Pair{i, j, 0})
				comCan = append(comCan, dfs(i, j, canCompo))
				canCompo++
			}
		}
	}

	Poligono = make([]Point, 0)
	dfs2(si, sj, canCompo, entradas[matriz[si][sj]][0])
	fmt.Println(Poligono)

	suma := 0
	for k := 0; k < len(comPos); k++ {
		fmt.Println(comPos[k], comCan[k], k)
		//for l := 0; l < 4; l++ {
		point := Point{float64(comPos[k].j) + 0.5, float64(comPos[k].i) + 0.5}

		if checkInside(Poligono, len(Poligono), point) {
			fmt.Println("esta dentro")
			suma += comCan[k]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if distancias[i][j] == 1000000000 {
				fmt.Print(marca[i][j])
			} else {
				fmt.Print(string(matriz[i][j]))
			}
		}
		fmt.Println()
	}

	fmt.Println(suma)
}
