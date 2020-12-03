package main

import (
    "fmt"
    "adventofcode/2020/modules/readinput"
)

func main() {
    numbers := readinput.ReadInts("inputs/1/input.txt", "\n")

    for index_1, number_1 := range numbers {
        for index_2, number_2 := range numbers[index_1+1:] {
            for _, number_3 := range numbers[index_2+1:] {
                if number_1 + number_2 + number_3 == 2020 {
                    fmt.Printf("%d\n", (number_1 * number_2 * number_3));

                    return
                }
            }
        }
    }
}