package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error connecting to server: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		_, err := fmt.Fprintln(conn, text)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error sending command: %v\n", err)
			os.Exit(1)
		}

		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading response: %v\n", err)
			os.Exit(1)
		}

		fmt.Print(response)
	}
}
