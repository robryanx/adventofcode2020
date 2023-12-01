package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2020/modules/readinput"
)

func main() {
	parts, err := readinput.ReadStrings("inputs/13/input.txt", "\n")
	if err != nil {
		panic(err)
	}

	bus_list := [][2]int{}
	n_list := []int{}
	a_list := []int{}
	N := 1

	for position, trip_time_string := range strings.Split(parts[1], ",") {
		if trip_time_string != "x" {
			id_number, err := strconv.Atoi(trip_time_string)
			if err != nil {
				panic(err)
			}

			record := [2]int{position, id_number}

			bus_list = append(bus_list, record)

			n_list = append(n_list, id_number)
			a_list = append(a_list, id_number-position)
			N *= id_number
		}
	}

	increment_method(bus_list)
	chinese_remainder_method(n_list, N, a_list)
}

func increment_method(bus_list [][2]int) {
	list_length := len(bus_list)

	counter := 0
	increment := bus_list[0][1]
	start := 1

	for {
		found := true
		for i := start; i < list_length; i++ {
			if (counter+bus_list[i][0])%bus_list[i][1] != 0 {
				found = false
				break
			} else {
				increment *= bus_list[i][1]
				start++
			}
		}

		if found {
			break
		}

		counter += increment
	}

	fmt.Println(counter)
}

func chinese_remainder_method(n_list []int, N int, a_list []int) {
	fmt.Println(chinese_remainder(n_list, N, a_list))
}

func chinese_remainder(n []int, N int, a []int) int {
	result := 0

	for i := range n {
		ai := a[i]
		ni := n[i]
		bi := int(math.Floor(float64(N) / float64(ni)))

		result += ai * bi * mod_inverse(bi, ni)
	}

	return result % N
}

func mod_inverse(a int, m int) int {
	return int(new(big.Int).ModInverse(big.NewInt(int64(a)), big.NewInt(int64(m))).Int64())
}
