package main

import (
	"seven/utils"
)

func main() {

	lines := utils.OpenFile("sample.txt")

	for _, l := range lines {
		utils.ProcessLine(l)
	}

}
