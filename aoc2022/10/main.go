package main

import (
	"aoc10/10/utils"
	"fmt"
	"strings"
)

var CheckPoints [6]int = [6]int{20, 60, 100, 140, 180, 220}

func main() {

	text := utils.ReadFile()

	numCycle := 0

	for _, value := range text {

		if value == "noop" {

			numCycle++

		} else if strings.HasPrefix(value, "addx") {

			fmt.Println(value)

			numCycle += 2

		}
	}

	fmt.Println(numCycle)

}
