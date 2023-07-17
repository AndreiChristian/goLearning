package main

import (
	"aoc/utils"
	"fmt"
)

func main() {

	X := 0
	Y := 0

	aim := 0

	text := utils.Scan("input.txt")

	for _, v := range text {

		command, value := utils.SplitText(v)

		switch command {
		case "forward":
			X = X + value
			Y = Y + value*aim
		case "down":
			aim = aim + value
		case "up":
			aim = aim - value

		}

	}

	fmt.Println(X * Y)

}
