package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Thanks for rating. Your rating:", input)

	numRating, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error", err)
		panic(err)
	}

	fmt.Printf("The parsed input is of type %T", numRating)

}
