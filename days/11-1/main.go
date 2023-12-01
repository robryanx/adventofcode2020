package main

import (
	"fmt"
	"strings"

	"github.com/robryanx/adventofcode2020/modules/readinput"
)

func main() {
	var seats [][]string

	lines, err := readinput.ReadStrings("inputs/11/input.txt", "\n")
	if err != nil {
		panic(err)
	}

	for i, line := range lines {
		line_seats := strings.Split(line, "")
		seats = append(seats, make([]string, len(line_seats)))

		for j, count := 0, len(line_seats); j < count; j++ {
			seats[i][j] = line_seats[j]
		}
	}

	var changed bool
	var occupied int
	for {
		changed, occupied, seats = run_process(seats)

		if !changed {
			break
		}
	}

	fmt.Println(occupied)
}

func run_process(seats [][]string) (bool, int, [][]string) {
	max_y := len(seats)
	max_x := len(seats[0])

	new_map := [][]string{}
	occupied := 0
	changed := false

	for y := 0; y < max_y; y++ {
		new_map = append(new_map, make([]string, max_x))

		for x := 0; x < max_x; x++ {
			new_map[y][x] = seats[y][x]

			if seats[y][x] != "." {
				filled := adjacent_filled(seats, x, y, max_x, max_y)

				if seats[y][x] == "L" && filled == 0 {
					new_map[y][x] = "#"
					changed = true
				} else if seats[y][x] == "#" && filled > 3 {
					new_map[y][x] = "L"
					changed = true
				}

				if new_map[y][x] == "#" {
					occupied++
				}
			}
		}
	}

	return changed, occupied, new_map
}

func adjacent_filled(seats [][]string, x int, y int, max_x int, max_y int) int {
	matrix := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
	}

	filled := 0
	for _, pair := range matrix {
		// check for bounds
		check_y := y + pair[0]
		check_x := x + pair[1]

		if check_y > -1 && check_y < max_y && check_x > -1 && check_x < max_x {
			if seats[check_y][check_x] == "#" {
				filled++
			}
		}
	}

	return filled
}
