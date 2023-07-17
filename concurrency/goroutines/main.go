package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	evilNinjas := []string{"Tom", "Bob", "John", "Mark"}

	for _, n := range evilNinjas {

		// like this we attack them one by one
		//  attack(n)

		go attack(n)

	}

	time.Sleep(time.Second / 2)
}

func attack(target string) {

	fmt.Println("Throwing ninja star at", target)
	time.Sleep(time.Second)

}
