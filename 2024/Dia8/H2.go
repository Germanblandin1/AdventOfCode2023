package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

var n int
var m int

var matrix [][]rune
var mapa map[rune][]Point
var visited [][]bool

func main() {
	var total uint64 = 0
	tam := 12
	tam = 50
	reader := bufio.NewReader(os.Stdin)
	matrix = make([][]rune, tam)
	mapa = make(map[rune][]Point)
	visited = make([][]bool, tam)
	n = tam
	for c := 0; c < tam; c++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		matrix[c] = []rune(line)
		//fmt.Println(matrix[c])
		visited[c] = make([]bool, len(line))
		for i := 0; i < len(line); i++ {
			visited[c][i] = false
			if line[i] == '.' {
				continue
			}
			_, ok := mapa[matrix[c][i]]
			if !ok {
				mapa[matrix[c][i]] = make([]Point, 0)
			}
			mapa[matrix[c][i]] = append(mapa[matrix[c][i]], Point{c, i})
		}
	}
	m = len(matrix[0])

	for _, v := range mapa {
		//fmt.Printf("rune %v tiene %v\n", string(r), len(v))
		if len(v) == 1 {
			continue
		}
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				//fmt.Printf("--Comparando %v %v\n", v[i], v[j])
				difx := v[j].x - v[i].x
				dify := v[j].y - v[i].y
				//fmt.Printf("---difx %v dify %v\n", difx, dify)

				//fmt.Printf("---apx %v apy %v\n", apx, apy)
				apx := v[i].x
				apy := v[i].y
				for {
					if !(apx < 0 || apx >= n || apy < 0 || apy >= m) {
						if !visited[apx][apy] {
							total++
							visited[apx][apy] = true
						}
					} else {
						break
					}
					apx = apx - difx
					apy = apy - dify

				}

				bpx := v[j].x
				bpy := v[j].y
				for {
					if !(bpx < 0 || bpx >= n || bpy < 0 || bpy >= m) {
						if !visited[bpx][bpy] {
							total++
							visited[bpx][bpy] = true
						}
					} else {
						break
					}
					bpx = bpx + difx
					bpy = bpy + dify

				}
			}
		}

	}

	fmt.Println(total)
}
