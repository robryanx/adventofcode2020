package main

import (
	"fmt"
	"strings"

	"github.com/robryanx/adventofcode2020/modules/readinput"
)

func main() {
	var display [][]bool

	lines, err := readinput.ReadStrings(3, false, "\n")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		var row []bool

		for _, char := range strings.Split(line, "") {
			row = append(row, (char == "#"))
		}

		display = append(display, row)
	}

	pos_x := 0
	pos_y := 0

	width := len(display[0])
	height := len(display)

	trees := 0
	for (pos_y + 1) < height {
		pos_x += 3
		pos_y++

		if display[pos_y][pos_x%width] {
			trees++
		}
	}

	fmt.Println(trees)
}
