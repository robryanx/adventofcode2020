package main

import (
	"fmt"
	"strings"

	"github.com/robryanx/adventofcode2020/modules/readinput"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	total := 0

	lines, err := readinput.ReadStrings(6, false, "\n\n")
	if err != nil {
		panic(err)
	}

	for _, groups := range lines {
		answers_collect := make(map[rune]int)

		answers := strings.Split(groups, "\n")
		answers_length := len(answers)

		for _, answer := range answers {
			for _, char := range answer {
				if _, ok := answers_collect[char]; ok {
					answers_collect[char]++
				} else {
					answers_collect[char] = 1
				}
			}
		}

		for _, count := range answers_collect {
			if count == answers_length {
				total++
			}
		}
	}

	fmt.Println(total)
}
