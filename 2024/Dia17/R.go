package main

import (
	"fmt"
	"math"
	"strings"
)

const (
	A = 4
	B = 5
	C = 6
)

var registros []int = []int{0, 1, 2, 3, 0, 0, 0}
var instrucciones []func(opr int)
var pointer int = 0
var output []string

func runProgram(program []int) {
	for pointer < len(program) {
		code := program[pointer]
		opr := program[pointer+1]
		instrucciones[code](opr)
		if code == 3 && registros[A] != 0 {
			continue
		}
		pointer += 2
	}

	fmt.Println(strings.Join(output, ","))
}

func main() {
	pointer = 0
	output = make([]string, 0)
	instrucciones = make([]func(opr int), 8)

	//adv
	instrucciones[0] = func(opr int) {
		num := registros[A]
		dem := math.Pow(2, float64(registros[opr]))
		registros[A] = int(num / int(dem))
	}

	//bxl
	instrucciones[1] = func(opr int) {
		registros[B] = registros[B] ^ opr
	}

	//bst
	instrucciones[2] = func(opr int) {
		registros[B] = registros[opr] % 8
	}

	//jnz
	instrucciones[3] = func(opr int) {
		if registros[A] != 0 {
			pointer = opr
		}
	}

	//bxc
	instrucciones[4] = func(opr int) {
		registros[B] = registros[B] ^ registros[C]
	}

	//out
	instrucciones[5] = func(opr int) {
		val := registros[opr] % 8
		strVal := string(rune(val) + '0')
		output = append(output, strVal)
	}

	//bdv
	instrucciones[6] = func(opr int) {
		num := registros[A]
		dem := math.Pow(2, float64(registros[opr]))
		registros[B] = int(num / int(dem))
	}

	//cdv
	instrucciones[7] = func(opr int) {
		num := registros[A]
		dem := math.Pow(2, float64(registros[opr]))
		registros[C] = int(num / int(dem))
	}

	//var total uint64 = 0

	var programStr string
	fmt.Scanf("Register A: %d\n", &registros[A])
	fmt.Scanf("Register B: %d\n", &registros[B])
	fmt.Scanf("Register C: %d\n", &registros[C])
	fmt.Scanf("Program: %s\n", &programStr)

	splitProgram := strings.Split(programStr, ",")
	var program []int = make([]int, len(splitProgram))

	for i, v := range splitProgram {
		program[i] = int(v[0] - '0')
	}
	//fmt.Println(program)
	runProgram(program)

	//fmt.Println(total)
}
