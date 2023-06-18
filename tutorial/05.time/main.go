package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time in Go")

	presentTime := time.Now()

	//! the string MUST BE formatted by very strict standards
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

}
