package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2020/util/readinput"
)

type Instruction struct {
	instruction_type  string
	instruction_value int
}

func main() {
	instruction_list := []*Instruction{}

	lines, err := readinput.ReadStrings(8, false, "\n")
	if err != nil {
		panic(err)
	}

	for _, instructions := range lines {
		instruction_parts := strings.Split(instructions, " ")

		value, err := strconv.Atoi(instruction_parts[1])
		if err != nil {
			panic(err)
		}

		current_instruction := &Instruction{instruction_type: instruction_parts[0], instruction_value: value}

		instruction_list = append(instruction_list, current_instruction)
	}

	accumulator := 0
	instruction_visited := make(map[int]bool)
	instruction_pointer := 0

	valid, accumulator_end := instruction_branch(instruction_list, instruction_visited, instruction_pointer, accumulator, false)

	if valid {
		fmt.Println(accumulator_end)
	}
}

func instruction_branch(instruction_list []*Instruction, instruction_visited map[int]bool, instruction_pointer int, accumulator int, branched bool) (bool, int) {
	valid := false
	accumulator_end := 0
	instruction_count := len(instruction_list)

	for {
		if instruction_pointer == instruction_count {
			valid = true
			break
		}

		instruction := instruction_list[instruction_pointer]

		if _, ok := instruction_visited[instruction_pointer]; ok {
			break
		}

		if instruction.instruction_type == "jmp" {
			if !branched {
				instruction_list[instruction_pointer].instruction_type = "nop"

				visited_copy := make(map[int]bool)
				for key, value := range instruction_visited {
					visited_copy[key] = value
				}

				valid, accumulator_end = instruction_branch(instruction_list, visited_copy, instruction_pointer, accumulator, true)
				if valid {
					accumulator = accumulator_end
					break
				}
			}

			instruction_visited[instruction_pointer] = true
			instruction_pointer += instruction.instruction_value
		} else if instruction.instruction_type == "nop" {
			if !branched {
				instruction_list[instruction_pointer].instruction_type = "jmp"

				visited_copy := make(map[int]bool)
				for key, value := range instruction_visited {
					visited_copy[key] = value
				}

				valid, accumulator_end = instruction_branch(instruction_list, visited_copy, instruction_pointer, accumulator, true)
				if valid {
					accumulator = accumulator_end
					break
				}
			}

			instruction_visited[instruction_pointer] = true
			instruction_pointer++
		} else {
			accumulator += instruction.instruction_value

			instruction_visited[instruction_pointer] = true
			instruction_pointer++
		}
	}

	return valid, accumulator
}
