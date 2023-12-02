package main

import (
	"fmt"
	"strings"

	"github.com/robryanx/adventofcode2020/modules/readinput"
)

func main() {
	total := 0

	lines, err := readinput.ReadStrings(6, false, "\n\n")
	if err != nil {
		panic(err)
	}

	for _, groups := range lines {
		answers_collect := make(map[rune]bool)

		answers := strings.Split(groups, "\n")

		for _, answer := range answers {
			for _, char := range answer {
				answers_collect[char] = true
			}
		}

		total += len(answers_collect)
	}

	fmt.Println(total)
}
