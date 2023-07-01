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

	var sumsSlice []int

	groups := strings.Split(string(data), "\n\n")

	for _, group := range groups {

		sum := 0

		numbers := strings.Split(group, "\n")

		for _, number := range numbers {

			if number == "" {
				continue
			}

			num, err := strconv.Atoi(number)

			if err != nil {
				log.Fatal(err)
			}

			sum = sum + num

		}

		sumsSlice = append(sumsSlice, sum)

	}

	sorted := bubbleSort(sumsSlice)

	sorted = sorted[len(sorted)-3:]

	sum := 0

	for _, n := range sorted {
		sum = sum + n
	}

	fmt.Println(sorted)

	fmt.Println(sum)

}

func bubbleSort(arr []int) []int {
	n := len(arr)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}
	return arr
}
