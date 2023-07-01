package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	// read the data from the file

	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	score := 0

	backpacks := strings.Split(string(data), "\n")

	for i, b := range backpacks {

		if i%3 != 0 {
			continue
		}

		for _, char := range b {

			if strings.ContainsRune(backpacks[i+1], char) && strings.ContainsRune(backpacks[i+2], char) {

				if char >= 65 && char <= 90 {

					score = score + int(char) - 65 + 27

				} else if char >= 97 && char <= 122 {

					score = score + int(char) - 96

				}

				break

			}

		}

	}

	fmt.Println(score)

}
