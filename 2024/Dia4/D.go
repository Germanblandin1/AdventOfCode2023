package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var movi = []int{0, 0, 1, -1, 1, -1, 1, -1}
var movj = []int{1, -1, 0, 0, 1, -1, -1, 1}

func main() {
	var total int = 0
	tam := 140
	reader := bufio.NewReader(os.Stdin)
	sopa := make([]string, tam)
	for i := 0; i < tam; i++ {
		line, _ := reader.ReadString('\n')
		sopa[i] = strings.TrimSpace(line)
	}

	n := len(sopa)
	m := len(sopa[0])
	fmt.Println(n, m)
	word := "XMAS"
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if sopa[i][j] == 'X' {
				for k := 0; k < 8; k++ {
					ii := i
					jj := j
					found := true
					for l := 0; l < 4; l++ {
						ii = i + movi[k]*l
						jj = j + movj[k]*l
						if ii < 0 || ii >= n || jj < 0 || jj >= m {
							found = false
							break
						}
						if sopa[ii][jj] != word[l] {
							found = false
							break
						}
					}
					if found {
						total++
					}
				}
			}
		}
	}

	fmt.Println(total)
}
