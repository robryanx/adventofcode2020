package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2020/util/readinput"
)

func main() {
	list, err := readinput.ReadInts(15, false, ",")
	if err != nil {
		panic(err)
	}

	lastSeen := make(map[int]int, 3000000)

	var prev int
	var count int

	for i, number := range list {
		lastSeen[number] = i
		count++
	}

	for {
		curr := 0
		if _, ok := lastSeen[prev]; ok {
			curr = count - lastSeen[prev]
		}
		lastSeen[prev] = count
		count++

		if count == 30000000 {
			break
		}

		prev = curr
	}

	fmt.Println(prev)
}
