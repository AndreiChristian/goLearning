package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	defer func() {
		fmt.Println(time.Since(now))
	}()

	smokeSignal := make(chan bool)

	evilNinja := "Walter"

	go attack(evilNinja, smokeSignal)

	smokeSignal <- false

	fmt.Println(<-smokeSignal)

}

func attack(target string, channel chan bool) {

	fmt.Println("Attacking", target)

	time.Sleep(time.Second)

	channel <- true

}
