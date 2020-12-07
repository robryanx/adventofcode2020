package main

import (
    "fmt"
    "strings"
    "adventofcode/2020/modules/readinput"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    total := 0
    for _, groups := range readinput.ReadStrings("inputs/6/input.txt", "\n\n") {
        answers_collect := make(map[rune]bool)

        answers := strings.Split(groups, "\n")

        for _, answer := range answers {
            for _, char := range answer {
                answers_collect[char] = true
            }
        }

        total += len(answers_collect)
    }

    fmt.Println(total)
}
