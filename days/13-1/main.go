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
    parts := readinput.ReadStrings("inputs/13/input.txt", "\n")

    leave_time, err := strconv.Atoi(parts[0])
    check(err)
    closest := -1
    closest_id := -1
    var remainder int

    for _, trip_time_string := range strings.Split(parts[1], ",") {
        if trip_time_string != "x" {
            id_number, err := strconv.Atoi(trip_time_string)
            check(err)

            remainder = leave_time % id_number
            if remainder > 0 {
                remainder = id_number - remainder
            }

            if closest == -1 || remainder < closest {
                closest = remainder
                closest_id = id_number
            }
        }
    }

    fmt.Println(closest * closest_id)
}
