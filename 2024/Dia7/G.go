package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Key struct {
	sum uint64
	i   int
}

var values []uint64
var memo map[Key]bool
var visited map[Key]bool
var goal uint64
var n int

func dp(key Key) bool {
	if key.sum == goal && key.i == n {
		return true
	}
	if key.i >= n || key.sum > goal {
		return false
	}
	if visited[key] {
		return memo[key]
	}

	visited[key] = true
	yes := dp(Key{key.sum + values[key.i], key.i + 1})
	no := dp(Key{key.sum * values[key.i], key.i + 1})

	concat := strconv.FormatUint(key.sum, 10) + strconv.FormatUint(values[key.i], 10)
	concatNum, _ := strconv.ParseUint(concat, 10, 64)
	other := dp(Key{concatNum, key.i + 1})
	memo[key] = yes || no || other
	return memo[key]

}

func main() {
	var total uint64 = 0
	tam := 9
	tam = 850
	reader := bufio.NewReader(os.Stdin)
	for c := 0; c < tam; c++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		strValues := strings.Split(line, " ")

		goal, _ = strconv.ParseUint(strings.Trim(strValues[0], ":"), 10, 64)
		values = make([]uint64, len(strValues)-1)
		for i := 1; i < len(strValues); i++ {
			values[i-1], _ = strconv.ParseUint(strValues[i], 10, 64)
		}
		n = len(values)

		memo = make(map[Key]bool)
		visited = make(map[Key]bool)

		//fmt.Printf("goal %v values %v\n", goal, values)
		if dp(Key{values[0], 1}) {
			total += goal
		}

	}

	fmt.Println(total)
}
