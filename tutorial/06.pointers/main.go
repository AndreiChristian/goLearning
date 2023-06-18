package main

import "fmt"

func main() {

	number := 10
	newNumber := number

	number = 12

	fmt.Println("The number is", number)
	fmt.Println("The memory address is", &number)

	fmt.Println("The number is", newNumber)
	fmt.Println("The memory address is", &newNumber)

}
