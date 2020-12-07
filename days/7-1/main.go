package main

import (
    "fmt"
    "strings"
    "regexp"
    "adventofcode/2020/modules/readinput"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    bag_contains := make(map[string][]string)
    for _, rules := range readinput.ReadStrings("inputs/7/input.txt", "\n") {
        rule_parts := strings.Split(rules, " bags contain ")

        r, _ := regexp.Compile("([0-9]+)\\s([\\sa-z]+)\\sbag")

        for _, contents := range strings.Split(rule_parts[1], ",") {
            contents_parts := r.FindStringSubmatch(contents)

            if len(contents_parts) > 0 {
                bag_contains[contents_parts[2]] = append(bag_contains[contents_parts[2]], rule_parts[0])
            }
        }
    }

    collect_bags := map[string]bool{}
    collect_bags = bag_recurse(bag_contains, collect_bags, "shiny gold")

    fmt.Println(len(collect_bags))
}

func bag_recurse(bag_contains map[string][]string, collect_bags map[string]bool, current_bag string) map[string]bool {
    for i, count := 0, len(bag_contains[current_bag]); i<count; i++ {
        collect_bags[bag_contains[current_bag][i]] = true

        collect_bags = bag_recurse(bag_contains, collect_bags, bag_contains[current_bag][i])
    }

    return collect_bags
}
