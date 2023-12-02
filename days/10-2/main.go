package main

import (
	"fmt"
	"sort"

	"github.com/robryanx/adventofcode2020/modules/readinput"
)

func main() {
	input, err := readinput.ReadInts(10, false, "\n")
	if err != nil {
		panic(err)
	}

	sort.Ints(input)

	input = append([]int{0}, input...)
	input = append(input, input[len(input)-1]+3)

	groupings := [][]int{}
	current_group := []int{}
	current_jolt := 0

	for i, count := 0, len(input); i < count; i++ {
		if input[i]-current_jolt == 3 {
			if len(current_group) > 0 {
				groupings = append(groupings, current_group)
				current_group = nil
			}
		}

		current_group = append(current_group, input[i])
		current_jolt = input[i]
	}

	acc := 1
	for _, group := range groupings {
		start := []int{}
		arrangement_count := arrangements(group, start, 0)

		acc *= arrangement_count
	}

	fmt.Println(acc)
}

func arrangements(input []int, start []int, pos int) int {
	arrangement_count := 0
	var current_jolt int
	if len(start) > 0 {
		current_jolt = start[len(start)-1]
	} else {
		current_jolt = input[pos]
		pos++
	}

	var options []int

	for {
		options = nil
		for i, count := pos, len(input); i < count; i++ {
			if input[i]-current_jolt < 4 {
				options = append(options, input[i])
			} else {
				break
			}
		}

		if len(options) == 0 {
			arrangement_count++

			break
		}

		if len(options) > 1 {
			for count, option := range options {
				build_start := make([]int, len(start))
				copy(build_start, start)
				build_start = append(build_start, option)

				arrangement_count += arrangements(input, build_start, (pos + 1 + count))
			}

			break
		} else {
			start = append(start, options[0])
			current_jolt = options[0]
			pos++
		}
	}

	return arrangement_count
}
