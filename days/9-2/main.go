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
    input := readinput.ReadInts("inputs/9/input.txt", "\n")

    invalid_number := 29221323

    outer_loop:
    for i, count := 0, len(input); i<count; i++ {
        sum := 0
        num_range := []int{}
        for j := i; j < count; j++ {
            num_range = append(num_range, input[j])
            sum += input[j]
            if sum == invalid_number {
                sort.Ints(num_range)
                fmt.Println(num_range[len(num_range)-1]+num_range[0])
                break outer_loop
            } else if sum > invalid_number {
                break
            }
        }
    }
}