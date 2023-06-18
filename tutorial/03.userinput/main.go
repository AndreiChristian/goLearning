package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	input, _ := reader.ReadString('\n')

	fmt.Println("Thank you. Your rating was:", input)
	fmt.Printf("Type of your rating is %T", input)

	
}
