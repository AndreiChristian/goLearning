package main

import (
	"fmt"
	"seven/utils"
)

func main() {

	lines := utils.OpenFile("sample.txt")

	fmt.Println(lines)

}
