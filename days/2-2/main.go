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

	lines, err := readinput.ReadStrings(2, false, "\n")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		matches := r.FindStringSubmatch(line)

		pos_a, err := strconv.Atoi(matches[1])
		pos_a--
		if err != nil {
			panic(err)
		}

		pos_b, err := strconv.Atoi(matches[2])
		pos_b--
		if err != nil {
			panic(err)
		}

		first_match := string(matches[4][pos_a]) == matches[3]
		second_match := string(matches[4][pos_b]) == matches[3]

		if first_match != second_match {
			valid_count++
		}
	}

	fmt.Println(valid_count)
}
