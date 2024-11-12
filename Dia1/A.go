package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

var dicc = map[string]string{
	"0":     "0",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var array = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func EsunNumero(s string) (string, bool) {
	val, ok := dicc[s]
	if ok {
		return val, ok
	}
	return "", ok
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var input string
	var total int
	for {
		_, err := fmt.Scanln(&input)
		if err != nil && err == io.EOF {
			break
		}
		first := ""
		last := ""
		input2 := input
		for {
			for j, value := range array {
				//fmt.Println("value", value)
				if strings.HasPrefix(input, value) {
					numero, _ := EsunNumero(array[j])
					if first == "" {
						//fmt.Println(numero)
						first = numero
						break
					}
				}
			}
			if first != "" {
				break
			}
			input = input[1:]
		}
		//fmt.Println(input2)
		for {
			for j, value := range array {
				//fmt.Println("value", value)
				if strings.HasPrefix(reverse(input2), reverse(value)) {
					numero, _ := EsunNumero(array[j])
					if last == "" {
						//fmt.Println(numero)
						last = numero
						break
					}
				}

			}
			if last != "" {
				break
			}
			input2 = input2[:len(input2)-1]
		}

		sum := first + last
		val, _ := strconv.Atoi(sum)
		total = total + val

	}
	fmt.Println(total)

}
