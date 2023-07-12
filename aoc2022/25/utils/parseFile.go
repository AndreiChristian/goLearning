package utils

import (
	"bufio"
	"log"
	"os"
)

func ParseFile() []string {
	file, err := os.Open("yourfile.txt")
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	return txtlines
}
