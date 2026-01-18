package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2020/util/readinput"
)

type parseLocation int32

const (
	parseLocationFields parseLocation = iota
	parseLocationYourTicket
	parseLocationNearbyTickets
)

func main() {
	lines, err := readinput.ReadStrings(16, false, "\n")
	if err != nil {
		panic(err)
	}

	fields := [][][2]int{}

	errorRate := 0

	loc := parseLocationFields
	for _, line := range lines {
		switch loc {
		case parseLocationFields:
			if line != "" {
				fields = append(fields, parseField(line))
			} else {
				loc = parseLocationYourTicket
			}
		case parseLocationYourTicket:
			if line == "" {
				loc = parseLocationNearbyTickets
			}
		case parseLocationNearbyTickets:
			if line == "nearby tickets:" {
				continue
			}

			value := invalidField(line, fields)
			if value != -1 {
				errorRate += value
			}
		}
	}

	fmt.Println(errorRate)
}

func parseField(line string) [][2]int {
	nameParts := strings.Split(line, ": ")

	fieldRanges := [][2]int{}
	for rangeStr := range strings.SplitSeq(nameParts[1], " or ") {
		rangeParts := strings.Split(rangeStr, "-")

		start, err := strconv.Atoi(rangeParts[0])
		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(rangeParts[1])
		if err != nil {
			panic(err)
		}

		fieldRanges = append(fieldRanges, [2]int{start, end})
	}

	return fieldRanges
}

func invalidField(line string, fields [][][2]int) int {
	for valueStr := range strings.SplitSeq(line, ",") {
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			panic(err)
		}

		found := false
		for _, field := range fields {
			if (value >= field[0][0] && value <= field[0][1]) ||
				(value >= field[1][0] && value <= field[1][1]) {
				found = true
				break
			}
		}

		if !found {
			return value
		}
	}

	return -1
}
