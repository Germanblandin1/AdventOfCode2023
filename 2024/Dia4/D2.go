package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var movi = []int{1, -1, 1, -1}
var movj = []int{1, -1, -1, 1}

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
	XMAS := [][]string{
		{"M.M", ".A.", "S.S"},
		{"S.S", ".A.", "M.M"},
		{"M.S", ".A.", "M.S"},
		{"S.M", ".A.", "S.M"},
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if sopa[i][j] == 'A' {
				for k := 0; k < 4; k++ {
					ii := i
					jj := j
					found := true
					for l := 0; l < 4; l++ {
						ii = i + movi[l]
						jj = j + movj[l]
						if ii < 0 || ii >= n || jj < 0 || jj >= m {
							found = false
							break
						}
						if sopa[ii][jj] != XMAS[k][1+movi[l]][1+movj[l]] {
							found = false
							break
						}
					}
					if found {
						total++
						break
					}
				}
			}
		}
	}

	fmt.Println(total)
}
