package main

import (
	"fmt"
	"io"
	"sort"
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
	listB := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%d %d\n", &listA[i], &listB[i])
		if err != nil && err == io.EOF {
			break
		}
		fmt.Println(listA[i], listB[i])
	}

	sort.Ints(listA)
	sort.Ints(listB)
	for i := 0; i < n; i++ {
		total += Abs(listA[i] - listB[i])
	}

	fmt.Println(total)
}
