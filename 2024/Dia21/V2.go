package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Horizontal = iota
	Vertical
)

var INF int = 0x3f3f3f3f

var movi = []int{-1, 0, 1, 0}
var movj = []int{0, -1, 0, 1}

var tecladoNum [][]int = [][]int{
	{7, 8, 9},
	{4, 5, 6},
	{1, 2, 3},
	{-1, 0, 10},
}

var posXYTecladoNum []Pair = []Pair{
	{3, 1},
	{2, 0},
	{2, 1},
	{2, 2},
	{1, 0},
	{1, 1},
	{1, 2},
	{0, 0},
	{0, 1},
	{0, 2},
	{3, 2},
}

var shortesPaths [][]string = [][]string{
	{"", "^<", "^", "^>", "^^<", "^^", "^^>", "^^^<", "^^^", "^^^>", ">"},
	{">v", "", ">", ">>", "^", "^>", "^>>", "^^", "^^>", "^^>>", ">>v"},
	{"v", "<", "", ">", "^<", "^", "^>", "^^<", "^^", "^^>", "v>"},
	{"v<", "<<", "<", "", "^<<", "^<", "^", "^^<<", "^^<", "^^", "v"},
	{">vv", "v", "v>", "v>>", "", ">", ">>", "^", "^>", "^>>", ">>vv"},
	{"vv", "v<", "v", "v>", "<", "", ">", "^<", "^", "^>", "vv>"},
	{"vv<", "v<<", "v<", "v", "<<", "<", "", "^<<", "^<", "^", "vv"},
	{">vvv", "vv", "vv>", "vv>>", "v", "v>", "v>>", "", ">", ">>", ">>vvv"},
	{"vvv", "vv<", "vv", "vv>", "v<", "v", "v>", "<", "", ">", "vvv>"},
	{"vvv<", "vv<<", "vv<", "vv", "v<<", "v<", "v", "<<", "<", "", "vvv"},
	{"<", "^<<", "^<", "^", "^^<<", "^^<", "^^", "^^^<<", "^^^<", "^^^", ""},
}

var tecladoDir [][]int = [][]int{
	{-1, 0, 4},
	{1, 2, 3},
}

var posXYTecladoDir []Pair = []Pair{
	{0, 1},
	{1, 0},
	{1, 1},
	{1, 2},
	{0, 2},
}

type Pair struct {
	i, j int
}

var posiciones []int

var shortesTeclado [][]string = [][]string{
	{"A", "v<A", "vA", "v>A", ">A"},
	{">^A", "A", ">A", ">>A", ">>^A"},
	{"^A", "<A", "A", ">A", "^>A"},
	{"^<A", "<<A", "<A", "A", "^A"},
	{"<A", "v<<A", "v<A", "vA", "A"},
}

var mapita map[string]int = map[string]int{
	"^": 0,
	"<": 1,
	"v": 2,
	">": 3,
	"A": 4,
}

var mapitaNum map[string]int = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"A": 10,
}

type estado struct {
	c     string
	nivel int
}

var dp map[estado]uint64

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func pathWriter(off, dir int) []byte {
	var path []byte
	var c byte
	if dir == Horizontal {
		if off < 0 {
			c = '>'
		} else {
			c = '<'
		}
	} else {
		if off < 0 {
			c = 'v'
		} else {
			c = '^'
		}
	}
	for range Abs(off) {
		path = append(path, c)
	}
	return path
}

func shortestSeq(srcp, dstp int, isNumPad bool) string {

	var path []byte

	var src, dst Pair
	if isNumPad {
		src = posXYTecladoNum[srcp]
		dst = posXYTecladoNum[dstp]
	} else {
		src = posXYTecladoDir[srcp]
		dst = posXYTecladoDir[dstp]
	}

	dr := src.i - dst.i
	dc := src.j - dst.j

	movesV := pathWriter(dr, Vertical)
	movesH := pathWriter(dc, Horizontal)

	var onGap bool
	if isNumPad {
		onGap = (src.i == 3 && dst.j == 0) || (src.j == 0 && dst.i == 3)
	} else {
		onGap = (src.j == 0 && dst.i == 0) || (src.i == 0 && dst.j == 0)
	}

	goingLeft := dst.j < src.j

	if goingLeft != onGap {
		movesV, movesH = movesH, movesV
	}

	path = append(append([]byte{}, movesV...), movesH...)
	path = append(path, 'A')
	//fmt.Println(string(path))
	return string(path)
}

func contar(c string, nivel int) uint64 {
	//fmt.Println(c, nivel)

	if val, ok := dp[estado{c, nivel}]; ok {
		return val
	}

	if c == "" {
		return 0
	}
	if nivel == -1 {
		//fmt.Printf("%v", c)
		return uint64(len(c))
	}
	dp[estado{c, nivel}] = 0
	var value uint64 = 0
	posAnt := posiciones[nivel]
	for _, v := range c {
		posSig := mapita[string(v)]
		value += contar(shortestSeq(posAnt, posSig, false), nivel-1)
		posAnt = posSig
	}
	dp[estado{c, nivel}] = value
	return value
}

func main() {
	var total uint64 = 0
	tam := 5
	reader := bufio.NewReader(os.Stdin)
	niveles := 25
	posiciones = make([]int, niveles)
	for i := 0; i < niveles; i++ {
		posiciones[i] = 4
	}

	dp = make(map[estado]uint64)

	for c := 0; c < tam; c++ {
		posiciones = make([]int, niveles)
		for i := 0; i < niveles; i++ {
			posiciones[i] = 4
		}
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		finalCode := line

		fmt.Println(finalCode)
		finalCodeInt, _ := strconv.Atoi(finalCode[0 : len(finalCode)-1])

		inicial := ""
		inicio := 10
		for i := 0; i < len(finalCode); i++ {
			inicial += shortestSeq(inicio, mapitaNum[string(finalCode[i])], true)
			inicio = mapitaNum[string(finalCode[i])]
		}
		val := contar(inicial, niveles-1)
		fmt.Println()
		fmt.Println("--", val)
		fmt.Println("--", finalCodeInt)
		total += uint64(val) * uint64(finalCodeInt)
	}

	fmt.Println(total)
}
