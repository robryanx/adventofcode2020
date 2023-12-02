package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2020/modules/readinput"
)

func main() {
	list, err := readinput.ReadInts(15, false, ",")
	if err != nil {
		panic(err)
	}

	last_seen := make(map[int]int)

	var current_number int
	var prev_number int

	var count int

	for i, number := range list {
		last_seen[number] = i
		count++
	}

	prev_number = 0
	current_number = 3

	for {
		if _, ok := last_seen[prev_number]; ok {
			current_number = count - last_seen[prev_number]
		} else {
			current_number = 0
		}

		last_seen[prev_number] = count
		count++

		prev_number = current_number

		if count+1 == 30000000 {
			break
		}
	}

	fmt.Println(current_number)
}
