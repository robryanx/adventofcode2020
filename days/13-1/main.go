package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2020/util/readinput"
)

func main() {
	parts, err := readinput.ReadStrings(13, false, "\n")
	if err != nil {
		panic(err)
	}

	leave_time, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	closest := -1
	closest_id := -1
	var remainder int

	for _, trip_time_string := range strings.Split(parts[1], ",") {
		if trip_time_string != "x" {
			id_number, err := strconv.Atoi(trip_time_string)
			if err != nil {
				panic(err)
			}

			remainder = leave_time % id_number
			if remainder > 0 {
				remainder = id_number - remainder
			}

			if closest == -1 || remainder < closest {
				closest = remainder
				closest_id = id_number
			}
		}
	}

	fmt.Println(closest * closest_id)
}
