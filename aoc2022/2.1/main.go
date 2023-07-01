package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	tournament := strings.Split(string(data), "\n")

	score := 0

	for _, games := range tournament {

		if games == "" {
			continue
		}

		choices := strings.Split(games, " ")

		opponentChoice := choices[0]
		result := choices[1]

		switch opponentChoice {
		case "A":

			switch result {

			case "X":

				score += 3

			case "Y":

				score += 4

			case "Z":

				score += 8

			}

		case "B":

			switch result {

			case "X":

				score += 1

			case "Y":

				score += 5

			case "Z":

				score += 9

			}

		case "C":

			switch result {

			case "X":

				score += 2

			case "Y":

				score += 6

			case "Z":

				score += 7

			}

		}

	}

	fmt.Println(score)

}
