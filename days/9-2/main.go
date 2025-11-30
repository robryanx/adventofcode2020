package main

import (
	"fmt"
	"sort"

	"github.com/robryanx/adventofcode2020/util/readinput"
)

func main() {
	input, err := readinput.ReadInts(9, false, "\n")
	if err != nil {
		panic(err)
	}

	invalid_number := 29221323

outer_loop:
	for i, count := 0, len(input); i < count; i++ {
		sum := 0
		num_range := []int{}
		for j := i; j < count; j++ {
			num_range = append(num_range, input[j])
			sum += input[j]
			if sum == invalid_number {
				sort.Ints(num_range)
				fmt.Println(num_range[len(num_range)-1] + num_range[0])
				break outer_loop
			} else if sum > invalid_number {
				break
			}
		}
	}
}
