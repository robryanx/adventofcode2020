package main

import (
    "fmt"
    "strings"
    "strconv"
    "sort"
    "adventofcode/2020/modules/readinput"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    var seats []uint32

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

        seats = append(seats, seat_id)
    }

    sort.Slice(seats, func(i, j int) bool { return seats[i] < seats[j] })

    for i:=1; i<=len(seats)-1; i++ {
        if (seats[i-1]+2) == seats[i] {
            fmt.Println(seats[i]-1)
        }
    }
}
