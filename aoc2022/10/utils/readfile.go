package utils

import (
	"bufio"
	"log"
	"os"
)

type Regiser struct {
}

var CheckPoints [6]int = [6]int{20, 60, 100, 140, 180, 220}

func ReadFile() []string {

	file, err := os.Open("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}

func Check(cycleNum int) bool {

	for _, value := range CheckPoints {

		if value == cycleNum {
			return true
		}

	}

	return false

}
