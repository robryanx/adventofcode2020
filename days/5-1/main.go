package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2020/modules/readinput"
)

func main() {
	biggest_seat := uint32(0)

	lines, err := readinput.ReadStrings(5, false, "\n")
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

		if seat_id > biggest_seat {
			biggest_seat = seat_id
		}
	}

	fmt.Println(biggest_seat)
}
