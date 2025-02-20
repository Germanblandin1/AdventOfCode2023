package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func stringsToInt(numbersStr []string) []int {
	numbers := make([]int, 0)
	for _, num := range numbersStr {
		val, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Error al convertir número:", err)
			return nil
		}
		numbers = append(numbers, val)
	}
	return numbers
}

func esSeguroArray(numbers []int) (bool, int) {
	fmt.Println("Array", numbers)
	n := len(numbers)
	var esSeguro bool = true
	var esDecreciente bool = true
	for j := 1; j < n; j++ {
		if j == 1 && numbers[j] > numbers[j-1] {
			esDecreciente = false
		}
		diff := numbers[j] - numbers[j-1]
		if diff > 3 || diff < -3 {
			esSeguro = false
			return false, j - 1
		}
		if esDecreciente && diff >= 0 {
			esSeguro = false
			return false, j - 1
		}
		if !esDecreciente && diff <= 0 {
			esSeguro = false
			return false, j - 1
		}

	}
	return esSeguro, -1
}

func main() {
	var total int = 0
	tam := 1000
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < tam; i++ {
		line, _ := reader.ReadString('\n')
		numbersStr := strings.Fields(line)
		numbers := stringsToInt(numbersStr)
		esSeguro, posProblema := esSeguroArray(numbers)
		if esSeguro {
			total++
		} else {
			var esSeguroA, esSeguroB, esSeguroC bool = false, false, false
			if posProblema-1 >= 0 {
				A := make([]int, len(numbers[:posProblema-1])+len(numbers[posProblema:]))
				copy(A, numbers[:posProblema-1])
				copy(A[len(numbers[:posProblema-1]):], numbers[posProblema:])
				esSeguroA, _ = esSeguroArray(A)
				fmt.Println("Array A:", A)
			}

			B := make([]int, len(numbers[:posProblema])+len(numbers[posProblema+1:]))
			copy(B, numbers[:posProblema])
			copy(B[len(numbers[:posProblema]):], numbers[posProblema+1:])
			esSeguroB, _ = esSeguroArray(B)
			fmt.Println("Array B:", B)

			// Array C (sin posProblema+1)
			if posProblema+1 < len(numbers) {
				C := make([]int, len(numbers[:posProblema+1])+len(numbers[posProblema+2:]))
				copy(C, numbers[:posProblema+1])
				copy(C[len(numbers[:posProblema+1]):], numbers[posProblema+2:])
				esSeguroC, _ = esSeguroArray(C)
				fmt.Println("Array C:", C)
			}
			if esSeguroA || esSeguroB || esSeguroC {
				total++
			}
		}

	}

	fmt.Println(total)
}
