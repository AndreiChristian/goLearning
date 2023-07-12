package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	quantity int
	from     int
	to       int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	stacks := make([][]string, 0)
	moves := make([]Move, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "move") {
			// This line describes a move
			parts := strings.Split(line, " ")
			quantity, _ := strconv.Atoi(parts[1])
			from, _ := strconv.Atoi(parts[3])
			to, _ := strconv.Atoi(parts[5])
			moves = append(moves, Move{quantity, from - 1, to - 1})
		} else if len(line) > 0 {
			// This line describes a stack
			stack := strings.Fields(line)
			if len(stack) == 1 {
				stacks = append(stacks, []string{})
			}
			stacks[len(stacks)-1] = append([]string{stack[0]}, stacks[len(stacks)-1]...)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, move := range moves {
		performMove(stacks, move)
	}

	result := make([]string, len(stacks))
	for i, stack := range stacks {
		result[i] = stack[0]
	}

	fmt.Println(strings.Join(result, ""))
}

func performMove(stacks [][]string, move Move) {
	for i := 0; i < move.quantity; i++ {
		// Pop from the 'from' stack
		crate := stacks[move.from][0]
		stacks[move.from] = stacks[move.from][1:]

		// Push onto the 'to' stack
		stacks[move.to] = append([]string{crate}, stacks[move.to]...)
	}
}
