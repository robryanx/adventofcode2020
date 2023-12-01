package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2020/modules/console"
	"github.com/robryanx/adventofcode2020/modules/readinput"
)

func main() {
	instruction_list := []*console.Instruction{}

	lines, err := readinput.ReadStrings("inputs/8/input.txt", "\n")
	if err != nil {
		panic(err)
	}

	for _, instructions := range lines {
		instruction_parts := strings.Split(instructions, " ")

		value, err := strconv.Atoi(instruction_parts[1])
		if err != nil {
			panic(err)
		}

		current_instruction := &console.Instruction{Instruction_type: instruction_parts[0], Instruction_value: value}

		instruction_list = append(instruction_list, current_instruction)
	}

	_, accumulator := console.Run(instruction_list)

	fmt.Println(accumulator)
}
