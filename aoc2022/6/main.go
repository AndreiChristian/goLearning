package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	file, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	text := strings.Split(string(file), "")

	i := 0

	for i+14 < len(text) {

		slice := text[i : i+14]

		window := make(map[string]int)

		for _, value := range slice {

			_, exists := window[value]

			if !exists {
				window[value] = 1

			}
		}

		if len(window) == 14 {
			log.Fatal(i + 14)
		}

		i++

	}

}
