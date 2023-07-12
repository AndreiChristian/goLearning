package main

import (
	"aoc/snafu"
	"aoc/utils"
	"fmt"
)

func main() {

	input := utils.ParseFile("sample.txt")

	snafu := snafu.New()

	snafu.SetMap()

	sum := 0

	for _, l := range input {
		values := utils.SplitLine(l)
		s, _ := snafu.ToInteger(values)
		fmt.Println(s)
		sum = sum + s
	}

	fmt.Println(sum)

	fmt.Println("")

	fmt.Println(snafu.ToInteger(utils.SplitLine("1=-0-2")))

}
