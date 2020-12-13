package main

import (
    "fmt"
    "sort"
    "adventofcode/2020/modules/readinput"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    input := readinput.ReadInts("inputs/10/input.txt", "\n")

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