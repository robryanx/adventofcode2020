package main

import (
    "fmt"
    "strings"
    "strconv"
    "adventofcode/2020/modules/readinput"
    "adventofcode/2020/modules/console"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    instruction_list := []*console.Instruction{}

    for _, instructions := range readinput.ReadStrings("inputs/8/input.txt", "\n") {
        instruction_parts := strings.Split(instructions, " ")

        value, err := strconv.Atoi(instruction_parts[1])
        check(err)

        current_instruction := &console.Instruction{Instruction_type: instruction_parts[0], Instruction_value: value}

        instruction_list = append(instruction_list, current_instruction)
    }

    _, accumulator := console.Run(instruction_list)

    fmt.Println(accumulator)
}