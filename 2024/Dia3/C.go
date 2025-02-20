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
	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	matches := re.FindAllStringSubmatch(string(content), -1)
	for _, match := range matches {
		a := match[1]
		b := match[2]
		ai, _ := strconv.ParseUint(a, 10, 64)
		bi, _ := strconv.ParseUint(b, 10, 64)
		total += ai * bi
	}
	fmt.Println(total)

}
