package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/robryanx/adventofcode2020/util/readinput"
)

func main() {
	// mask = 1001X0X00110011X01X1000110100011000X
	r_mask, _ := regexp.Compile("mask\\s\\=\\s([10X]{36})")

	// mem[9042] = 12448
	r_mem, _ := regexp.Compile("mem\\[([0-9]+)\\]\\s=\\s([0-9]+)")

	current_mask := ""
	memory := map[int]int64{}

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
				pos, err := strconv.Atoi(mem_matches[1])
				if err != nil {
					panic(err)
				}

				raw_val, err := strconv.ParseInt(mem_matches[2], 10, 64)
				if err != nil {
					panic(err)
				}

				val := apply_mask(raw_val, current_mask)

				memory[pos] = val
			}
		}
	}

	var total int64
	for _, val := range memory {
		total += val
	}

	fmt.Println(total)
}

func apply_mask(value int64, mask string) int64 {
	for i := 35; i >= 0; i-- {
		if mask[i] == '1' {
			value = setBit(value, uint(35-i))
		} else if mask[i] == '0' {
			value = clearBit(value, uint(35-i))
		}
	}

	return value
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
