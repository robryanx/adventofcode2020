package main

import (
	"fmt"
	"sort"

	"github.com/robryanx/adventofcode2020/util/readinput"
)

func main() {
	input, err := readinput.ReadInts(10, false, "\n")
	if err != nil {
		panic(err)
	}

	sort.Ints(input)

	difference_count := [3]int{0, 0, 0}

	current_jolt := 0
	for _, adaptor := range input {
		difference := adaptor - current_jolt

		difference_count[difference-1]++

		current_jolt = adaptor
	}

	difference_count[2]++

	fmt.Println(difference_count[0] * difference_count[2])
}
