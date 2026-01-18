package main

import (
	"fmt"
	"slices"
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
	myTicket := []int{}
	var combinedCandidates [][]int

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
			if line == "your ticket:" {
				continue
			}

			if line == "" {
				loc = parseLocationNearbyTickets
				continue
			}

			myTicket = ticketValues(line)
		case parseLocationNearbyTickets:
			if line == "nearby tickets:" {
				continue
			}

			valid, candidates := fieldCandidates(line, fields)
			if valid {
				if len(combinedCandidates) == 0 {
					combinedCandidates = candidates
				} else {
					for i := 0; i < len(combinedCandidates); i++ {
						newCandidates := []int{}
						for j := 0; j < len(combinedCandidates[i]); j++ {
							for k := 0; k < len(candidates[i]); k++ {
								if combinedCandidates[i][j] == candidates[i][k] {
									newCandidates = append(newCandidates, candidates[i][k])
									break
								}
							}
						}
						combinedCandidates[i] = newCandidates
					}
				}
			}
		}
	}

	queue := []int{}
	head := 0

	for i := 0; i < len(combinedCandidates); i++ {
		if len(combinedCandidates[i]) == 1 {
			queue = append(queue, i)
		}
	}

	for head < len(queue) {
		curr := queue[head]
		head++

		for i := len(combinedCandidates) - 1; i >= 0; i-- {
			if i != curr {
				for j := len(combinedCandidates[i]) - 1; j >= 0; j-- {
					if combinedCandidates[i][j] == combinedCandidates[curr][0] {
						combinedCandidates[i] = slices.Delete(combinedCandidates[i], j, j+1)
						if len(combinedCandidates[i]) == 1 {
							queue = append(queue, i)
						}
						break
					}
				}
			}
		}
	}

	ticketValue := 1
	for i, candidate := range combinedCandidates {
		if candidate[0] < 6 {
			ticketValue *= myTicket[i]
		}
	}

	fmt.Println(ticketValue)
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

func fieldCandidates(line string, fields [][][2]int) (bool, [][]int) {
	candidates := [][]int{}
	for valueStr := range strings.SplitSeq(line, ",") {
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			panic(err)
		}

		fieldCandidate := []int{}
		for i, field := range fields {
			if (value >= field[0][0] && value <= field[0][1]) ||
				(value >= field[1][0] && value <= field[1][1]) {
				fieldCandidate = append(fieldCandidate, i)
			}
		}

		if len(fieldCandidate) == 0 {
			return false, nil
		}

		candidates = append(candidates, fieldCandidate)
	}

	return true, candidates
}

func ticketValues(line string) []int {
	values := []int{}
	for valueStr := range strings.SplitSeq(line, ",") {
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			panic(err)
		}

		values = append(values, value)
	}

	return values
}
