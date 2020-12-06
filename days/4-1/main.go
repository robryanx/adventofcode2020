package main

import (
    "fmt"
    "strings"
    "adventofcode/2020/modules/readinput"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    required_fields := []string{
        "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",
    }

    valid_passports := 0
    for _, passport_raw := range readinput.ReadStrings("inputs/4/input.txt", "\n\n") {
        passport := parse_passport(passport_raw)

        valid := true
        for _, required_field := range required_fields {
            if _, ok := passport[required_field]; !ok {
                valid = false
                break
            }
        }

        if valid {
            valid_passports++;
        }
    }

    fmt.Println(valid_passports)
}

func parse_passport(passport_raw string) map[string]string {
    passport := make(map[string]string)

    for _, line := range strings.Split(passport_raw, "\n") {
        for _, field := range strings.Fields(line) {
            field_parts := strings.Split(field, ":")

            passport[field_parts[0]] = field_parts[1]
        }
    }

    return passport
}