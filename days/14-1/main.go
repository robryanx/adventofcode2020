package main

import (
    "fmt"
    "strings"
    "strconv"
    "adventofcode/2020/modules/readinput"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    current_mask := ""
    memory := map[int]int{}
    for _, row := range readinput.ReadStrings("inputs/14/input.txt", "\n") {
        // check to see if the row is a value of a mask
    }
}
