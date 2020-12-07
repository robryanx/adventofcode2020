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
        answers_collect := make(map[rune]int)

        answers := strings.Split(groups, "\n")
        answers_length := len(answers)

        for _, answer := range answers {
            for _, char := range answer {
                if _, ok := answers_collect[char]; ok {
                    answers_collect[char]++
                } else {
                    answers_collect[char] = 1
                }
            }
        }

        for _, count := range answers_collect {
            if count == answers_length {
                total++
            }
        }
    }

    fmt.Println(total)
}
