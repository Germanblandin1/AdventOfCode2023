package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

var basura string
var t int
var n, m int
var INF int = 1000000000

type Point struct {
	X, Y, Z float64
}

type Line struct {
	P1, P2 Point
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
		// Collinear
		return 0
	} else if val < 0 {
		// Anti-clockwise direction
		return 2
	}
	// Clockwise direction
	return 1
}

func calculateIntersection(l1, l2 Line) Point {
	a1 := l1.P2.Y - l1.P1.Y
	b1 := l1.P1.X - l1.P2.X
	c1 := a1*l1.P1.X + b1*l1.P1.Y

	a2 := l2.P2.Y - l2.P1.Y
	b2 := l2.P1.X - l2.P2.X
	c2 := a2*l2.P1.X + b2*l2.P1.Y

	determinant := a1*b2 - a2*b1

	if determinant == 0 {
		// The lines are parallel. This is simplified
		// by returning a pair of FLT_MAX
		return Point{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64}
	} else {
		x := (b2*c1 - b1*c2) / determinant
		y := (a1*c2 - a2*c1) / determinant
		return Point{x, y, 0}
	}
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

var lineas []Line

func checkInside(p Point, boundary Line) bool {
	if p.X < boundary.P1.X || p.X > boundary.P2.X || p.Y < boundary.P1.Y || p.Y > boundary.P2.Y {
		return false
	}
	return true
}

func buscarPunto(p Point, velocity Point, boundary Line) Point {
	li := float64(1)
	ls := 400000000000001.0
	for ls-li > float64(0.01) {

		mid := (li + ls) / float64(2)
		//fmt.Println(li, ls, mid)
		newPoint := Point{p.X + mid*velocity.X, p.Y + mid*velocity.Y, p.Z + mid*velocity.Z}
		if checkInside(newPoint, boundary) {
			li = mid
		} else {
			ls = mid
		}
	}
	return Point{p.X + li*velocity.X, p.Y + li*velocity.Y, p.Z + li*velocity.Z}
}

var maxi, mini float64

func main() {
	fmt.Scan(&n)
	fmt.Scan(&mini, &maxi)
	//box := Line{Point{mini, mini, mini}, Point{maxi, maxi, maxi}}

	lineas = make([]Line, n)
	//boundary := Line{Point{0, 0, 0}, Point{100, 100, 100}}
	for i := 0; i < n; i++ {
		var x, y, z, vx, vy, vz float64
		fmt.Scan(&x, &y, &z, &vx, &vy, &vz)
		lineas[i] = Line{Point{x, y, z}, Point{vx, vy, vz}}
		// if x < boundary.P1.X {
		// 	boundary.P1.X = x
		// }
		// if y < boundary.P1.Y {
		// 	boundary.P1.Y = y
		// }
		// if z < boundary.P1.Z {
		// 	boundary.P1.Z = z
		// }
		// if x > boundary.P2.X {
		// 	boundary.P2.X = x
		// }
		// if y > boundary.P2.Y {
		// 	boundary.P2.Y = y
		// }
		// if z > boundary.P2.Z {
		// 	boundary.P2.Z = z
		// }
	}
	// for i := 0; i < n; i++ {
	// 	//fmt.Println(boundary)
	// 	nextPoint := buscarPunto(lineas[i].P1, lineas[i].P2, boundary)

	// 	lineas[i].P2 = nextPoint
	// }

	// suma := 0
	// for i := 0; i < n; i++ {
	// 	fmt.Println(lineas[i].P1.X, lineas[i].P1.Y, lineas[i].P1.Z, lineas[i].P2.X, lineas[i].P2.Y, lineas[i].P2.Z)
	// 	for j := i + 1; j < n; j++ {
	// 		if isIntersect(lineas[i], lineas[j]) {
	// 			//fmt.Println(lineas[i].P1.X, lineas[i].P1.Y, lineas[i].P1.Z, lineas[i].P2.X, lineas[i].P2.Y, lineas[i].P2.Z)
	// 			//fmt.Println(lineas[j].P1.X, lineas[j].P1.Y, lineas[j].P1.Z, lineas[j].P2.X, lineas[j].P2.Y, lineas[j].P2.Z)

	// 			punto := calculateIntersection(lineas[i], lineas[j])
	// 			if checkInside(punto, box) {
	// 				suma++
	// 			}
	// 		}
	// 	}
	// }
	//fmt.Println(suma)

	matriz := make([]float64, 16)
	bres := make([]float64, 4)

	j := 0
	for i := 0; i < 4; i++ {
		matriz[j] = lineas[i].P2.Y - lineas[i+1].P2.Y
		j++
		matriz[j] = lineas[i+1].P2.X - lineas[i].P2.X
		j++
		matriz[j] = lineas[i+1].P1.Y - lineas[i].P1.Y
		j++
		matriz[j] = lineas[i].P1.X - lineas[i+1].P1.X
		j++
		bres[i] = lineas[i].P1.X*lineas[i].P2.Y - lineas[i].P1.Y*lineas[i].P2.X - lineas[i+1].P1.X*lineas[i+1].P2.Y + lineas[i+1].P1.Y*lineas[i+1].P2.X
	}
	a := mat.NewDense(4, 4, matriz)
	b := mat.NewVecDense(4, bres)
	x := mat.NewVecDense(4, nil)
	err := x.SolveVec(a, b)
	if err != nil {
		fmt.Println(err)
	}

	posX := x.AtVec(0)
	posY := x.AtVec(1)

	j = 0
	for i := 0; i < 4; i++ {
		matriz[j] = lineas[i].P2.Z - lineas[i+1].P2.Z
		j++
		matriz[j] = lineas[i+1].P2.X - lineas[i].P2.X
		j++
		matriz[j] = lineas[i+1].P1.Z - lineas[i].P1.Z
		j++
		matriz[j] = lineas[i].P1.X - lineas[i+1].P1.X
		j++
		bres[i] = lineas[i].P1.X*lineas[i].P2.Z - lineas[i].P1.Z*lineas[i].P2.X - lineas[i+1].P1.X*lineas[i+1].P2.Z + lineas[i+1].P1.Z*lineas[i+1].P2.X
	}

	a = mat.NewDense(4, 4, matriz)
	b = mat.NewVecDense(4, bres)
	x = mat.NewVecDense(4, nil)
	err = x.SolveVec(a, b)
	if err != nil {
		fmt.Println(err)
	}
	posZ := x.AtVec(1)
	fmt.Println(posX, posY, posZ, posX+posY+posZ)
	fmt.Printf("%0.6f", posX+posY+posZ)

}
