package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var patterns []string
var n int
var pattern string

var dpvalues []uint64
var mark []bool

func dp(pos int) uint64 {
	//fmt.Println(pos)
	//fmt.Println(pattern[pos:])
	if pos == n {
		return 1
	}
	if pos > n {
		return 0
	}
	if mark[pos] {
		return dpvalues[pos]
	}
	mark[pos] = true
	dpvalues[pos] = 0
	for i := 0; i < len(patterns); i++ {
		//fmt.Println("--", patterns[i])
		//fmt.Println("--", strings.HasPrefix(pattern[pos:], patterns[i]))
		if strings.HasPrefix(pattern[pos:], patterns[i]) {
			dpvalues[pos] += dp(pos + len(patterns[i]))
		}
	}
	return dpvalues[pos]
}

func clean() {
	dpvalues = make([]uint64, n)
	mark = make([]bool, n)
}

func main() {
	var total uint64 = 0
	tam := 8
	tam = 400
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	patrs := strings.Split(line, " ")
	patterns = make([]string, len(patrs))
	for i := 0; i < len(patrs); i++ {
		patterns[i] = strings.Trim(patrs[i], ",")
	}

	total = 0
	for c := 0; c < tam; c++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		n = len(line)
		pattern = line
		clean()
		total += dp(0)
	}

	fmt.Println(total)
}
