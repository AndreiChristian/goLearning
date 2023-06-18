package main

import "fmt"

type User struct {
	Name     string
	Email    string
	verified bool
	age      int
}

func main() {

	andrei := User{"andrei", "andrei@email.com", true, 12}

	fmt.Println(andrei.age)

	if num := 6; num < 5 {
		fmt.Println(num)
	}

}
