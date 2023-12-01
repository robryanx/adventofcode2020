package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2020/modules/readinput"
)

func main() {
	numbers, err := readinput.ReadInts("inputs/1/input.txt", "\n")
	if err != nil {
		panic(err)
	}

	for index_1, number_1 := range numbers {
		for _, number_2 := range numbers[index_1+1:] {
			if number_1+number_2 == 2020 {
				fmt.Printf("%d\n", (number_1 * number_2))

				return
			}
		}
	}
}
