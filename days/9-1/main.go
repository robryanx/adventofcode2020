package main

import (
    "fmt"
    "adventofcode/2020/modules/readinput"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    input := readinput.ReadInts("inputs/9/input.txt", "\n")

    for i, count := 25, len(input); i<count; i++ {
        check_numbers := input[i-25:i]

        valid := false
        for j := 0; j < 25; j++ {
            for k := (j+1); k < 25; k++ {
                if check_numbers[j] + check_numbers[k] == input[i] {
                    valid = true
                    break
                }
            }
        }

        if !valid {
            fmt.Println(input[i])
            break
        }
    }
}