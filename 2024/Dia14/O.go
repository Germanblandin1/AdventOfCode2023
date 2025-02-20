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

func main() {

	var total uint64 = 0
	tam := 12
	tam = 500
	n := 101
	m := 103
	veces := 100
	//reader := bufio.NewReader(os.Stdin)
	c1, c2, c3, c4 := 0, 0, 0, 0
	for c := 0; c < tam; c++ {
		var x, y, vx, vy int

		fmt.Scanf("p=%d,%d v=%d,%d\n", &x, &y, &vx, &vy)

		nx := modAdd(x, modMul(veces, vx, n), n)
		ny := modAdd(y, modMul(veces, vy, m), m)
		//fmt.Println(nx, ny)
		nmid := n / 2
		mmid := m / 2

		//revisar los 4 cuadrantes
		if nx < nmid && ny < mmid {
			c1++
		} else if nx < nmid && ny > mmid {
			c2++
		} else if nx > nmid && ny < mmid {
			c3++
		} else if nx > nmid && ny > mmid {
			c4++
		}
	}
	total = uint64(c1 * c2 * c3 * c4)
	fmt.Println(total)
}
