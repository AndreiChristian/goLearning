package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Scan(fileName string) []string {

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}

func SplitText(s string) (string, int) {

	splitted := strings.Split(s, " ")

	command := splitted[0]

	value, err := strconv.Atoi(splitted[1])

	if err != nil {
		log.Fatal(err)
	}

	return command, value

}
