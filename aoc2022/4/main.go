package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {

	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	score := 0

	pairs := strings.Split(string(data), "\n")

	for _, pair := range pairs {

		if pair == "" {
			continue
		}

		pairSlice := strings.Split(pair, ",")

		first := strings.Split(pairSlice[0], "-")
		second := strings.Split(pairSlice[1], "-")

		firstStart, _ := strconv.Atoi(first[0])
		firstEnd, _ := strconv.Atoi(first[1])
		secondStart, _ := strconv.Atoi(second[0])
		secondEnd, _ := strconv.Atoi(second[1])

		if firstStart <= secondEnd && firstEnd >= secondStart {
			score += 1
		}
	}

	fmt.Println(score)
}
