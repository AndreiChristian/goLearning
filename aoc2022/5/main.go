package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	// create the stacks

	data, err := ioutil.ReadFile("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	// we split by the first empty line
	// the first element will be the crates
	// the second element will be the instructions
	info := strings.Split(string(data), "\n\n")

	cratesMap := make(map[string][]string)

	cratesInfo := info[0]
	cratesLines := strings.Split(cratesInfo, "\n")
	cratesNumbersLine := cratesLines[len(cratesLines)-1]
	crateNumbers := strings.Split(cratesNumbersLine, " ")

	for _, crateNumber := range crateNumbers {

		cratesMap[crateNumber] = []string{}

	}

	// crates := cratesLines[:len(cratesLines)-1]

	fmt.Println(cratesLines)

}
