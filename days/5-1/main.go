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
    biggest_seat := uint32(0)
    for _, pass := range readinput.ReadStrings("inputs/5/input.txt", "\n") {
        runes := []rune(pass)

        row := string(runes[0:7])
        col := string(runes[7:])

        row = strings.ReplaceAll(row, "F", "0")
        row = strings.ReplaceAll(row, "B", "1")

        col = strings.ReplaceAll(col, "L", "0")
        col = strings.ReplaceAll(col, "R", "1")

        row_num, err := strconv.ParseInt(row, 2, 32)
        check(err)

        col_num, err := strconv.ParseInt(col, 2, 32)
        check(err)

        seat_id := uint32(row_num * 8 + col_num)

        if seat_id > biggest_seat {
            biggest_seat = seat_id
        }
    }

    fmt.Println(biggest_seat)
}
