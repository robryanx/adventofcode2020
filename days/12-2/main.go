package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/robryanx/adventofcode2020/modules/readinput"
)

func main() {
	direction_value := map[string][]int{
		"E": {1, 0},
		"S": {0, -1},
		"W": {-1, 0},
		"N": {0, 1},
	}

	position := [2]int{0, 0}
	waypoint := [2]int{10, 1}
	new_waypoint := [2]int{0, 0}
	var waypoint_angle float64

	lines, err := readinput.ReadStrings(12, false, "\n")
	if err != nil {
		panic(err)
	}

	for _, instruction := range lines {
		direction := string(instruction[0])
		movement_value, err := strconv.Atoi(instruction[1:])
		if err != nil {
			panic(err)
		}

		switch direction {
		case "N", "S", "E", "W":
			waypoint[0] += direction_value[direction][0] * movement_value
			waypoint[1] += direction_value[direction][1] * movement_value
		case "L", "R":
			if direction == "R" {
				waypoint_angle = -float64(movement_value)
			} else if direction == "L" {
				waypoint_angle = float64(movement_value)
			}

			waypoint_angle_radians := waypoint_angle * math.Pi / 180

			new_waypoint[0] = int(math.Round(math.Cos(waypoint_angle_radians)*float64(waypoint[0]) - math.Sin(waypoint_angle_radians)*float64(waypoint[1])))
			new_waypoint[1] = int(math.Round(math.Sin(waypoint_angle_radians)*float64(waypoint[0]) + math.Cos(waypoint_angle_radians)*float64(waypoint[1])))

			waypoint = new_waypoint
		case "F":
			position[0] += waypoint[0] * movement_value
			position[1] += waypoint[1] * movement_value
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
