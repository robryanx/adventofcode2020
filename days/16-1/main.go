package main

import (
	"github.com/robryanx/adventofcode2020/modules/readinput"
)

func main() {
	_, err := readinput.ReadInts(15, false, ",")
	if err != nil {
		panic(err)
	}
}
