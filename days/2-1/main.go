package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/robryanx/adventofcode2020/modules/readinput"
)

var r = regexp.MustCompile(`([0-9]+)\-([0-9]+)\s([a-z]+)\:\s(.*)`)

func main() {
	valid_count := 0

	lines, err := readinput.ReadStrings("inputs/2/input.txt", "\n")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		matches := r.FindStringSubmatch(line)

		min, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}

		max, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}

		letter_count := 0
		for _, letter := range matches[4] {
			if string(letter) == matches[3] {
				letter_count++
			}
		}

		if letter_count >= min && letter_count <= max {
			valid_count++
		}
	}

	fmt.Println(valid_count)
}
