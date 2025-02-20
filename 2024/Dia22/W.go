package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var n int
var secretNumbers []int64
var pruneValue int64 = 16777216

func nextSecretNumber(number int64) int64 {
	number = ((number * 64) ^ number) % pruneValue
	number = ((number / 32) ^ number) % pruneValue
	number = ((number * 2048) ^ number) % pruneValue
	return number
}

func main() {
	var total uint64 = 0
	tam := 4
	tam = 1749
	n = tam
	veces := 2000
	secretNumbers = make([]int64, tam)
	reader := bufio.NewReader(os.Stdin)
	for c := 0; c < tam; c++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		secretNumbers[c], _ = strconv.ParseInt(line, 10, 64)
	}

	for i := 0; i < n; i++ {
		original := secretNumbers[i]
		for j := 0; j < veces; j++ {
			secretNumbers[i] = nextSecretNumber(secretNumbers[i])
			//fmt.Println(original, secretNumbers[i])
		}
		fmt.Println(original, secretNumbers[i])
		fmt.Println()
		total += uint64(secretNumbers[i])
	}

	fmt.Println(total)
}
