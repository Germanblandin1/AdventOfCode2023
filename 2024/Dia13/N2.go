package main

import (
	"fmt"
)

func main() {

	var total uint64 = 0
	tam := 16
	tam = 1280
	casos := tam / 4

	//reader := bufio.NewReader(os.Stdin)
	for c := 0; c < casos; c++ {
		var A, B, C, D, E, F int64

		fmt.Scanf("Button A: X+%d, Y+%d\n", &A, &D)
		fmt.Scanf("Button B: X+%d, Y+%d\n", &B, &E)
		fmt.Scanf("Prize: X=%d, Y=%d\n\n", &C, &F)
		C += 10000000000000
		F += 10000000000000
		fmt.Println(A, B, C, D, E, F)

		var X, Y int64
		var num, dem int64

		num = A*F - D*C
		dem = -D*B + A*E

		if dem == 0 {
			fmt.Println("No solution")
			fmt.Printf("A=%d, B=%d, C=%d, D=%d, E=%d, F=%d\n", A, B, C, D, E, F)
			continue
		}
		Y = num / dem
		X = (C - B*Y) / A

		if X < 0 || Y < 0 || num%dem != 0 {
			continue
		}
		fmt.Println("X:", X, "Y:", Y)
		total += uint64(X*3 + Y*1)
	}

	fmt.Println(total)
}
