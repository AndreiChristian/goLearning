package utils

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func OpenFile(fileName string) []string {

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines

}

func ProcessLine(line string) error {

	if len(line) == 0 {
		return errors.New("the line is blank")
	}

	words := strings.Split(line, " ")

	initialSign := words[0]

	switch initialSign {
	case "$":
		fmt.Println("It is a command ", initialSign)

		switch words[1] {
		case "ls":

			// we will have some until the next command only files, so we should create

		case "cd":
			//
			switch words[2] {
			case "..":
				// move up
			case "/":
				// move back to the file
			default:
				// we move in a file

			}

		}

	case "dir":
		fmt.Println("It is a dir ", initialSign)
	default:
		fmt.Println("It is a file ", initialSign)

		// if it is a file, we have the name and the size, so we can create a node based on the parent, set its name and size

	}

	return nil

}
