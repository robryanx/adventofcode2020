package main

import (
	"github.com/robryanx/adventofcode2020/util/readinput"
)

func main() {
	_, err := readinput.ReadInts(15, false, ",")
	if err != nil {
		panic(err)
	}
}
