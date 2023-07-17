package utils

import (
	"bufio"
	"log"
	"os"
)

type Regiser struct {
}

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

func CheckForBenchMark(m map[int]int, cycleNumber int, sum int, registerValue int) {

	_, ok := m[cycleNumber]

	if !ok {
		return
	}

	sum = sum + registerValue*cycleNumber

}
