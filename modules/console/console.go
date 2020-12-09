package console

type Instruction struct {
    Instruction_type string
    Instruction_value int
}

func Run(instruction_list []*Instruction) (bool, int) {
	instruction_pointer := 0
	accumulator := 0
	instruction_visited := make(map[int]bool)

	instruction_count := len(instruction_list)

	valid := false

	for {
		if instruction_pointer == instruction_count {
		    valid = true
		    break
		}

	    if _, ok := instruction_visited[instruction_pointer]; ok {
	        break
	    } else {
	        instruction_visited[instruction_pointer] = true
	    }

	    instruction := instruction_list[instruction_pointer]

	    switch instruction.Instruction_type {
	        case "jmp":
	        	instruction_pointer += instruction.Instruction_value
	        	break;
	        case "acc":
	        	accumulator += instruction.Instruction_value

	        	instruction_pointer++
	        	break;
	        case "nop":
	        	instruction_pointer++
	        	break;
	    }
	}

	return valid, accumulator
}