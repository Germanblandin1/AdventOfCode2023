package main

import (
	"fmt"
)

// modAdd performs modular addition: (a + b) % n
func modAdd(a, b, n int) int {
	return ((a % n) + (b % n) + n) % n
}

// modMul performs modular multiplication: (a * b) % n
// Handles negative values of b or v correctly.
func modMul(a, b, n int) int {
	result := (a % n) * (b % n)
	if result < 0 {
		result += n
	}
	return result % n
}

var final [][]rune

func generateXmasTree(n, m, tam int) {
	final = make([][]rune, n)
	for i := 0; i < n; i++ {
		final[i] = make([]rune, m)
		for j := 0; j < m; j++ {
			final[i][j] = '.'
		}
	}

	iniI := (n - (tam + 16)) / 2

	nivel := 1
	iniJ := (m-(tam))/2 + (tam / 2) - 1
	for i := iniI; i < iniI+tam; i++ {
		for j := iniJ; j < iniJ+nivel; j++ {
			final[i][j] = 'X'
		}
		nivel += 2
		iniJ--
	}
	base := iniI + tam

	for i := 0; i < 16; i++ {
		final[base+i][m/2] = 'X'
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%c", final[i][j])
		}
		fmt.Println()
	}

}

type Pair struct {
	x, y int
}

type Robot struct {
	pos Pair
	vel Pair
}

var robots []Robot

func cleanFinal(n, m int) {
	final = make([][]rune, n)
	for i := 0; i < n; i++ {
		final[i] = make([]rune, m)
		for j := 0; j < m; j++ {
			final[i][j] = '.'
		}
	}
}

func PrintFinal(n, m int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%c", final[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func main() {

	//generateXmasTree(103, 101, 22)

	var total uint64 = 0
	tam := 12
	tam = 500
	n := 101
	m := 103
	//reader := bufio.NewReader(os.Stdin)
	robots = make([]Robot, tam)
	cleanFinal(m, n)
	for c := 0; c < tam; c++ {
		var x, y, vx, vy int

		fmt.Scanf("p=%d,%d v=%d,%d\n", &x, &y, &vx, &vy)

		robots[c] = Robot{Pair{x, y}, Pair{vx, vy}}
	}

	veces := 0
	for {
		if veces == 100000 {
			break
		}
		for i := 0; i < tam; i++ {
			final[robots[i].pos.y][robots[i].pos.x] = '.'
			robots[i].pos.x = modAdd(robots[i].pos.x, robots[i].vel.x, n)
			robots[i].pos.y = modAdd(robots[i].pos.y, robots[i].vel.y, m)
			final[robots[i].pos.y][robots[i].pos.x] = 'X'

		}
		fmt.Println(veces)
		PrintFinal(m, n)

		veces++
	}
	total = uint64(veces)
	fmt.Println(total)
	//6620

}
