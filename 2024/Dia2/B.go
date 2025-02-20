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
	numbers := make([]int, len(numbersStr))
	for i, num := range numbersStr {
		val, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Error al convertir número:", err)
			return nil
		}
		numbers[i] = val
	}
	return numbers
}

func main() {
	var total int = 0
	tam := 1000
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < tam; i++ {
		line, _ := reader.ReadString('\n')
		numbersStr := strings.Fields(line)
		numbers := stringsToInt(numbersStr)
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
				break
			}
			if esDecreciente && diff >= 0 {
				esSeguro = false
				break
			}
			if !esDecreciente && diff <= 0 {
				esSeguro = false
				break
			}

		}
		if esSeguro {
			total++
		}

	}

	fmt.Println(total)
}
