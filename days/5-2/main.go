package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2020/modules/readinput"
)

func main() {
	var seats []uint32

	lines, err := readinput.ReadStrings("inputs/5/input.txt", "\n")
	if err != nil {
		panic(err)
	}

	for _, pass := range lines {
		runes := []rune(pass)

		row := string(runes[0:7])
		col := string(runes[7:])

		row = strings.ReplaceAll(row, "F", "0")
		row = strings.ReplaceAll(row, "B", "1")

		col = strings.ReplaceAll(col, "L", "0")
		col = strings.ReplaceAll(col, "R", "1")

		row_num, err := strconv.ParseInt(row, 2, 32)
		if err != nil {
			panic(err)
		}

		col_num, err := strconv.ParseInt(col, 2, 32)
		if err != nil {
			panic(err)
		}

		seat_id := uint32(row_num*8 + col_num)

		seats = append(seats, seat_id)
	}

	sort.Slice(seats, func(i, j int) bool { return seats[i] < seats[j] })

	for i := 1; i <= len(seats)-1; i++ {
		if (seats[i-1] + 2) == seats[i] {
			fmt.Println(seats[i] - 1)
		}
	}
}
