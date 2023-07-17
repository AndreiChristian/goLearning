package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

func StringSliceToIntSlice(sa []string) []int {

	var intSlice []int

	for _, value := range sa {

		num, err := strconv.Atoi(value)

		if err != nil {
			log.Fatal(err)
		}

		intSlice = append(intSlice, num)

	}

	return intSlice

}
