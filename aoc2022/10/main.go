package main

import (
	"aoc10/10/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {

	text := utils.ReadFile()

	numCycle := 0
	sum := 0
	register := 1

	for _, value := range text {

		if value == "noop" {

			if utils.Check(numCycle) {

				fmt.Printf("Hit a check point: %v and the value is %v \n", numCycle, register)
				sum = sum + numCycle*register

			}

			numCycle++

		} else if strings.HasPrefix(value, "addx") {

			splitted := strings.Split(value, " ")

			registerChangeValue, err := strconv.Atoi(splitted[1])

			if err != nil {
				log.Fatal(err)

			}

			if utils.Check(numCycle) {

				fmt.Printf("Hit a check point: %v and the value is %v \n", numCycle, register)
				sum = sum + numCycle*register

			}

			numCycle++

			if utils.Check(numCycle) {

				fmt.Printf("Hit a check point: %v and the value is %v \n", numCycle, register)

				sum = sum + numCycle*register

			}

			numCycle++

			register = register + registerChangeValue

		}
	}

	fmt.Println()
	fmt.Println(sum)

}
