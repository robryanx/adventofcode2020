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

	width := len(display[0])
	height := len(display)

	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	total := 0
	for _, slope := range slopes {
		pos_x := 0
		pos_y := 0

		trees := 0
		for (pos_y + slope[1]) < height {
			pos_x += slope[0]
			pos_y += slope[1]

			if display[pos_y][pos_x%width] {
				trees++
			}
		}

		if total == 0 {
			total = trees
		} else {
			total *= trees
		}
	}

	fmt.Println(total)
}
