package main

import (
	"fmt"
	"keygo/internal/tcp"
	"log"
)

func main() {

	server, err := tcp.New("localhost:6379")

	if err != nil {

		log.Fatalf("error creating server : %v", err)

	}

	fmt.Println("Listenting on localhost:6379")

	err = server.Listen()

	if err != nil {

		log.Fatalf("error listening: %v", err)

	}

}
