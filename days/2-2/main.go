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

        pos_a, err := strconv.Atoi(matches[1])
        pos_a--
        check(err)

        pos_b, err := strconv.Atoi(matches[2])
        pos_b--
        check(err)

        first_match := string(matches[4][pos_a]) == matches[3]
        second_match := string(matches[4][pos_b]) == matches[3]

        if first_match != second_match {
            valid_count++
        }
    }

    fmt.Println(valid_count)
}