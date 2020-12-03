package main

import (
    "fmt"
    "regexp"
    "strconv"
    "adventofcode/2020/modules/readinput"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    r, _ := regexp.Compile("([0-9]+)\\-([0-9]+)\\s([a-z]+)\\:\\s(.*)")

    valid_count := 0
    for _, line := range readinput.ReadStrings("inputs/2/input.txt", "\n") {
        matches := r.FindStringSubmatch(line);

        min, err := strconv.Atoi(matches[1])
        check(err)

        max, err := strconv.Atoi(matches[2])
        check(err)

        letter_count := 0
        for _, letter := range matches[4] {
            if string(letter) == matches[3] {
                letter_count++
            }
        }

        if letter_count >= min && letter_count <= max {
            valid_count++
        }
    }

    fmt.Println(valid_count)
}