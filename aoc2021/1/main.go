package main

import (
	"aoc1/utils"
	"fmt"
)

func main() {

	text := utils.Scan("input.txt")

	nums := utils.StringSliceToIntSlice(text)

	baseWindow := nums[0] + nums[1] + nums[2]

	count := 0

	i := 1

	for i < len(nums)-2 {

		newWindow := nums[i] + nums[i+1] + nums[i+2]

		if newWindow > baseWindow {
			count++
		}
		i++

		baseWindow = newWindow

	}

	fmt.Println(count)

}
