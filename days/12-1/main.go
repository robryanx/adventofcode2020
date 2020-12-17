package main

import (
    "fmt"
    "strconv"
    "adventofcode/2020/modules/readinput"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    direction_list := []string{
       "E", "S", "W", "N",
    }

    direction_value := map[string][]int{
       "E": {1, 0},
       "S": {0, -1},
       "W": {-1, 0},
       "N": {0, 1},
    }

    facing_value := 0
    facing_pos := 0
    position := [2]int{0, 0}
    for _, instruction := range readinput.ReadStrings("inputs/12/input.txt", "\n") {
        direction := string(instruction[0])
        movement_value, err := strconv.Atoi(instruction[1:])
        check(err)

        switch direction {
        case "N", "S", "E", "W":
            position[0] += direction_value[direction][0] * movement_value
            position[1] += direction_value[direction][1] * movement_value
        case "L":
            facing_value -= (movement_value/90)
        case "R":
            facing_value += (movement_value/90)
        case "F":
            position[0] += direction_value[direction_list[facing_pos]][0] * movement_value
            position[1] += direction_value[direction_list[facing_pos]][1] * movement_value
        }

        facing_pos = facing_value % 4
        if facing_pos < 0 {
            facing_pos = 4 + facing_pos
        }
    }

    fmt.Println(Abs(position[0]) + Abs(position[1]))
}

func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}