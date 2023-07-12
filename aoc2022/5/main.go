package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	input, err := ioutil.ReadFile("./sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	parsedInput := string(input)

	splitedInput := strings.Split(parsedInput, "\n\n")

	var crates string
	var commands string

	if len(splitedInput) != 2 {
		log.Fatal("Error parsing the file")
	} else {

		crates = splitedInput[0]
		commands = splitedInput[1]

	}

	fmt.Println(crates)
	fmt.Println("")
	fmt.Println(commands)

}
