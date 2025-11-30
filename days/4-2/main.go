package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2020/util/readinput"
)

func main() {
	valid_passports := 0

	lines, err := readinput.ReadStrings(4, false, "\n\n")
	if err != nil {
		panic(err)
	}

	for _, passport_raw := range lines {
		passport := parse_passport(passport_raw)

		if validate_passport(passport) {
			valid_passports++
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

func validate_passport(passport map[string]string) bool {
	required_fields := []string{
		"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",
	}

	// check that all required fields are present
	for _, required_field := range required_fields {
		if _, ok := passport[required_field]; !ok {
			return false
		}
	}

	// (Birth Year) - four digits; at least 1920 and at most 2002.
	byr, err := strconv.Atoi(passport["byr"])
	if err != nil {
		return false
	}

	if byr < 1920 || byr > 2002 {
		return false
	}

	// (Issue Year) - four digits; at least 2010 and at most 2020.
	iyr, err := strconv.Atoi(passport["iyr"])
	if err != nil {
		return false
	}

	if iyr < 2010 || iyr > 2020 {
		return false
	}

	// (Expiration Year) - four digits; at least 2020 and at most 2030.
	eyr, err := strconv.Atoi(passport["eyr"])
	if err != nil {
		return false
	}

	if eyr < 2020 || eyr > 2030 {
		return false
	}

	// (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	r := regexp.MustCompile("([0-9]+)(cm|in)")
	matches := r.FindStringSubmatch(passport["hgt"])

	if len(matches) != 3 {
		return false
	} else {
		hgt, err := strconv.Atoi(matches[1])
		if err != nil {
			return false
		}

		if matches[2] == "cm" {
			if hgt < 150 || hgt > 193 {
				return false
			}
		} else {
			if hgt < 59 || hgt > 76 {
				return false
			}
		}
	}

	// (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	r_hair := regexp.MustCompile("^#([0-9a-f]+)$")
	matches = r_hair.FindStringSubmatch(passport["hcl"])

	if len(matches) != 2 {
		return false
	}

	eye_colors := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	// (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	if _, ok := eye_colors[passport["ecl"]]; !ok {
		return false
	}

	// (Passport ID) - a nine-digit number, including leading zeroes.
	r_passport := regexp.MustCompile("^([0-9]{9})$")
	matches = r_passport.FindStringSubmatch(passport["pid"])

	return len(matches) == 2
}
