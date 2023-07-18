package main

import "fmt"

func main(){

	channel := make(chan string,3)
	channel <- "First message"
	channel <- "Second message"
	channel <- "Third message"

	fmt.Println(<-channel)
	fmt.Println(<-channel)


}
