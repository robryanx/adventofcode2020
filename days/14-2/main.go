package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/robryanx/adventofcode2020/modules/readinput"
)

func main() {
	// mask = 1001X0X00110011X01X1000110100011000X
	r_mask, _ := regexp.Compile("mask\\s\\=\\s([10X]{36})")

	// mem[9042] = 12448
	r_mem, _ := regexp.Compile("mem\\[([0-9]+)\\]\\s=\\s([0-9]+)")

	current_mask := ""
	memory := map[int64]int{}

	lines, err := readinput.ReadStrings(14, false, "\n")
	if err != nil {
		panic(err)
	}

	for _, row := range lines {
		mask_matches := r_mask.FindStringSubmatch(row)
		if mask_matches != nil {
			current_mask = mask_matches[1]
		} else {
			mem_matches := r_mem.FindStringSubmatch(row)
			if mem_matches != nil {
				raw_pos, err := strconv.ParseInt(mem_matches[1], 10, 64)
				if err != nil {
					panic(err)
				}

				val, err := strconv.Atoi(mem_matches[2])
				if err != nil {
					panic(err)
				}

				positions := apply_mask(raw_pos, current_mask)
				for _, pos := range positions {
					memory[pos] = val
				}
			}
		}
	}

	var total int
	for _, val := range memory {
		total += val
	}

	fmt.Println(total)
}

func apply_mask(value int64, mask string) []int64 {
	floating_pos := []uint{}

	for i := 35; i >= 0; i-- {
		if mask[i] == '1' {
			value = setBit(value, uint(35-i))
		} else if mask[i] == 'X' {
			floating_pos = append(floating_pos, uint(35-i))
		}
	}

	values := apply_floating(value, floating_pos)

	return values
}

func apply_floating(value int64, floating_pos []uint) []int64 {
	values := []int64{}

	combinations := int(math.Pow(2, float64(len(floating_pos))))
	for i := 0; i < combinations; i++ {
		set_value := value
		for j := 0; j < len(floating_pos); j++ {
			if hasBit(i, uint(j)) {
				set_value = setBit(set_value, floating_pos[j])
			} else {
				set_value = clearBit(set_value, floating_pos[j])
			}
		}

		values = append(values, set_value)
	}

	return values
}

func apply_floating_recursive(value int64, floating_pos []uint) []int64 {
	values := []int64{}

	if len(floating_pos) > 0 {
		set_value := setBit(value, floating_pos[0])
		for i := 1; i < len(floating_pos); i++ {
			set_value = clearBit(set_value, floating_pos[i])
		}

		values = append(values, set_value)

		if len(floating_pos) > 1 {
			values = append(values, apply_floating(set_value, floating_pos[1:])...)
		}

		clear_value := clearBit(value, floating_pos[0])
		for i := 1; i < len(floating_pos); i++ {
			clear_value = clearBit(clear_value, floating_pos[i])
		}

		values = append(values, clear_value)

		if len(floating_pos) > 1 {
			values = append(values, apply_floating(clear_value, floating_pos[1:])...)
		}
	}

	return values
}

func setBit(n int64, pos uint) int64 {
	n |= (1 << pos)
	return n
}

func clearBit(n int64, pos uint) int64 {
	var mask int64
	mask = ^(1 << pos)
	n &= mask
	return n
}

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}
