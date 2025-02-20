package main

import (
	"fmt"
	"strconv"
	"strings"
)

var basura string
var t int
var n, m int
var INF int = 1000000000
var actual int

type lente struct {
	label string
	value int
}

var matriz []string
var cajas [][]lente

func main() {
	fmt.Scanln(&basura)
	cadenas := strings.Split(basura, ",")
	n = len(cadenas)
	total := 0
	cajas := make([][]lente, 256)
	for i := 0; i < 256; i++ {
		cajas[i] = make([]lente, 0)
	}
	for i := 0; i < n; i++ {
		m = len(cadenas[i])
		suma := 0
		divisor := ""
		for j := 0; j < m; j++ {
			if cadenas[i][j] == '=' || cadenas[i][j] == '-' {
				divisor = string(cadenas[i][j])
				break
			}
			suma += int(cadenas[i][j])
			suma *= 17
			suma %= 256
		}
		if divisor == "-" {
			label := strings.Split(cadenas[i], "-")[0]
			for j := 0; j < len(cajas[suma]); j++ {
				if label == cajas[suma][j].label {
					cajas[suma] = append(cajas[suma][:j], cajas[suma][j+1:]...)
					break
				}
			}
		} else if divisor == "=" {
			label, strvalue := strings.Split(cadenas[i], "=")[0], strings.Split(cadenas[i], "=")[1]
			value, _ := strconv.Atoi(strvalue)
			esta := false
			for j := 0; j < len(cajas[suma]); j++ {
				if label == cajas[suma][j].label {
					esta = true
					cajas[suma][j].value = value
					break
				}
			}
			if !esta {
				cajas[suma] = append(cajas[suma], lente{label, value})
			}
		}

	}
	for i := 0; i < 256; i++ {
		for j := 0; j < len(cajas[i]); j++ {
			total += (i + 1) * (j + 1) * cajas[i][j].value
		}
	}
	fmt.Println(total)
}
