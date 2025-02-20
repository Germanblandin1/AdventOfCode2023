package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	content, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("Error leyendo el archivo: %v\n", err)
		return
	}

	var total uint64 = 0
	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(string(content), -1)
	//fmt.Println(matches)
	var activo bool = true
	for _, match := range matches {
		if match[0] == "do()" {
			activo = true
			continue
		}
		if match[0] == "don't()" {
			activo = false
			continue
		}
		if activo {
			a := match[1]
			b := match[2]
			ai, _ := strconv.ParseUint(a, 10, 64)
			bi, _ := strconv.ParseUint(b, 10, 64)
			total += ai * bi
		}
	}
	fmt.Println(total)

}
