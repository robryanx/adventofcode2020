package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2020/modules/readinput"
)

func main() {
	list, err := readinput.ReadInts("inputs/15/input.txt", ",")
	if err != nil {
		panic(err)
	}

	list = append(list, 0)

	count := len(list)

	for {
		current_number := list[len(list)-1]
		last_seen := len(list)
		for i := len(list) - 2; i >= 0; i-- {
			if list[i] == current_number {
				last_seen = i
				break
			}
		}

		if last_seen == len(list) {
			list = append(list, 0)
		} else {
			list = append(list, len(list)-last_seen-1)
		}

		count++

		if count == 2020 {
			break
		}
	}

	fmt.Println(list[len(list)-1])
}
