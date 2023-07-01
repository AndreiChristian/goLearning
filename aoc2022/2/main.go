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
		myChoice := choices[1]

		switch myChoice {
		case "X": // Rock
			score += 1
			switch opponentChoice {
			case "A": // draw
				score += 3
			case "B": // lose, no extra points
			case "C": // win
				score += 6
			}
		case "Y": // Paper
			score += 2
			switch opponentChoice {
			case "A": // win
				score += 6
			case "B": // draw
				score += 3
			case "C": // lose, no extra points
			}
		case "Z": // Scissors
			score += 3
			switch opponentChoice {
			case "A": // lose, no extra points
			case "B": // win
				score += 6
			case "C": // draw
				score += 3
			}
		}
	}

	fmt.Println(score)

}
