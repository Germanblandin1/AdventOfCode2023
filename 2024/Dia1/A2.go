package main

import (
	"fmt"
	"io"
)

func Abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func main() {
	var total int = 0
	n := 1000
	listA := make([]int, n)
	listB := make(map[int]int, 0)
	var b int
	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%d %d\n", &listA[i], &b)
		if err != nil && err == io.EOF {
			break
		}
		listB[b]++
	}

	for i := 0; i < n; i++ {
		total += listA[i] * listB[listA[i]]
	}

	fmt.Println(total)
}
