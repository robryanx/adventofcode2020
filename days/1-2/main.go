package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2020/util/readinput"
)

func main() {
	numbers, err := readinput.ReadInts(1, false, "\n")
	if err != nil {
		panic(err)
	}

	for index_1, number_1 := range numbers {
		for index_2, number_2 := range numbers[index_1+1:] {
			for _, number_3 := range numbers[index_2+1:] {
				if number_1+number_2+number_3 == 2020 {
					fmt.Printf("%d\n", (number_1 * number_2 * number_3))

					return
				}
			}
		}
	}
}
