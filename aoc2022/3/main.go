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

	for _, b := range backpacks {

		first := b[:len(b)/2]
		second := b[len(b)/2:]

		for _, char := range first {

			if strings.ContainsRune(second, char) {

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
