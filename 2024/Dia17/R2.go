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

var registros []uint64 = []uint64{0, 1, 2, 3, 0, 0, 0}
var instrucciones []func(opr uint64)
var pointer uint64 = 0
var output []string
var splitProgram []string

var ant int = 0

func runProgram(program []uint64) string {
	for pointer < uint64(len(program)) {
		code := program[pointer]
		opr := program[pointer+1]
		instrucciones[code](opr)
		if code == 3 && registros[A] != 0 {
			continue
		}
		// if code == 5 {
		// 	if output[len(output)-1] != splitProgram[len(output)-1] {
		// 		break
		// 	}
		// 	if len(output) > len(splitProgram) {
		// 		break
		// 	}
		// }
		pointer += 2
	}

	return strings.Join(output, ",")
}

func clean() {
	registros[A] = 0
	registros[B] = 0
	registros[C] = 0
	pointer = 0
	output = make([]string, 0)
}

func f(val uint64, program []uint64, final string) bool {
	clean()
	registros[A] = val
	res := runProgram(program)
	if len(res) != ant {
		fmt.Println("Tam", len(res))
		ant = len(res)
	}
	fmt.Println(res)
	return res == final
}

var values = []uint64{1, 3, 1, 1, 1, 1, 1, 1}
var program []uint64
var final string

func Backtraking(i, pos int, intento uint64) uint64 {
	if i > 16 {
		return 0
	}

	result := f(intento, program, final)
	if result {
		return intento
	} else {
		for j := 1; j <= 8; j++ {
			nnewval := uint64(float64(j-1) * math.Pow(8, float64(15-i)))
			intento += nnewval

			f(intento, program, final)
			if output[pos] == splitProgram[pos] {
				res := Backtraking(i+1, pos-1, intento)
				if res != 0 {
					return res
				}
			}
			intento -= nnewval
		}
	}
	return 0
}

func main() {
	pointer = 0
	output = make([]string, 0)
	instrucciones = make([]func(opr uint64), 8)

	//adv
	instrucciones[0] = func(opr uint64) {
		num := registros[A]
		dem := math.Pow(2, float64(registros[opr]))
		registros[A] = uint64(num / uint64(dem))
	}

	//bxl
	instrucciones[1] = func(opr uint64) {
		registros[B] = registros[B] ^ opr
	}

	//bst
	instrucciones[2] = func(opr uint64) {
		registros[B] = registros[opr] % 8
	}

	//jnz
	instrucciones[3] = func(opr uint64) {
		if registros[A] != 0 {
			pointer = opr
		}
	}

	//bxc
	instrucciones[4] = func(opr uint64) {
		registros[B] = registros[B] ^ registros[C]
	}

	//out
	instrucciones[5] = func(opr uint64) {
		val := registros[opr] % 8
		strVal := string(rune(val) + '0')
		output = append(output, strVal)
	}

	//bdv
	instrucciones[6] = func(opr uint64) {
		num := registros[A]
		dem := uint64(1) << uint64(registros[opr])
		//dem := math.Pow(2, float64(registros[opr]))
		registros[B] = num / dem
	}

	//cdv
	instrucciones[7] = func(opr uint64) {
		num := registros[A]
		dem := math.Pow(2, float64(registros[opr]))
		registros[C] = uint64(num / uint64(dem))
	}

	//var total uint64 = 0

	var programStr string
	fmt.Scanf("Register A: %d\n", &registros[A])
	fmt.Scanf("Register B: %d\n", &registros[B])
	fmt.Scanf("Register C: %d\n", &registros[C])
	fmt.Scanf("Program: %s\n", &programStr)

	splitProgram = strings.Split(programStr, ",")
	//fmt.Println(splitProgram)
	program = make([]uint64, len(splitProgram))

	for i, v := range splitProgram {
		program[i] = uint64(v[0] - '0')
	}
	//fmt.Println(program)
	//runProgram(program)

	final = strings.TrimSpace(programStr)

	inicio := 1 * math.Pow(8, 15)

	fmt.Println(Backtraking(0, len(splitProgram)-1, uint64(inicio)))

	//fmt.Println(total)
}
