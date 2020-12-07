package main

import (
    "fmt"
    "strings"
    "strconv"
    "regexp"
    "adventofcode/2020/modules/readinput"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    bag_rules := make(map[string]map[string]int)
    for _, rules := range readinput.ReadStrings("inputs/7/input.txt", "\n") {
        rule_parts := strings.Split(rules, " bags contain ")

        r, _ := regexp.Compile("([0-9]+)\\s([\\sa-z]+)\\sbag")

        bag_rules[rule_parts[0]] = make(map[string]int)

        for _, contents := range strings.Split(rule_parts[1], ",") {
            contents_parts := r.FindStringSubmatch(contents)

            if len(contents_parts) > 0 {
                count, err := strconv.Atoi(contents_parts[1])
                check(err)

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
        total_count += bag_count + bag_count * bag_recurse(bag_contains, bag)
    }

    return total_count
}