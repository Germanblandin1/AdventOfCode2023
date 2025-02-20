package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var patterns []string

func main() {
	var total uint64 = 0
	tam := 8
	tam = 400
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	patrs := strings.Split(line, " ")
	patterns = make([]string, len(patrs))
	pattern := `^(`
	for i := 0; i < len(patrs); i++ {
		patterns[i] = strings.Trim(patrs[i], ",")
		pattern += patterns[i]
		if i < len(patrs)-1 {
			pattern += "|"
		}
	}
	pattern += ")+$"

	//pattern := `^(ABC|DDC|RRR)+$`
	regex := regexp.MustCompile(pattern)
	total = 0
	for c := 0; c < tam; c++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if regex.MatchString(line) {
			total++
		}
	}

	fmt.Println(total)
}
