package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2020/modules/readinput"
)

func main() {
	bag_rules := make(map[string]map[string]int)

	lines, err := readinput.ReadStrings(7, false, "\n")
	if err != nil {
		panic(err)
	}

	for _, rules := range lines {
		rule_parts := strings.Split(rules, " bags contain ")

		r := regexp.MustCompile(`([0-9]+)\s([\sa-z]+)\sbag`)

		bag_rules[rule_parts[0]] = make(map[string]int)

		for _, contents := range strings.Split(rule_parts[1], ",") {
			contents_parts := r.FindStringSubmatch(contents)

			if len(contents_parts) > 0 {
				count, err := strconv.Atoi(contents_parts[1])
				if err != nil {
					panic(err)
				}

				bag_rules[rule_parts[0]][contents_parts[2]] = count
			}
		}
	}

	bag_count := bag_recurse(bag_rules, "shiny gold")

	fmt.Println(bag_count)
}

func bag_recurse(bag_contains map[string]map[string]int, current_bag string) int {
	total_count := 0

	for bag, bag_count := range bag_contains[current_bag] {
		total_count += bag_count + bag_count*bag_recurse(bag_contains, bag)
	}

	return total_count
}
