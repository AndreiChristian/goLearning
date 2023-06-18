package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitList = []string{"Apple", "Banana", "Peach"}
	// fmt.Printf("Type of %T", fruitList)
	fruitList = append(fruitList, "Lemon")
	fruitList = append(fruitList, "Orange")

	fruitList = append(fruitList[1:4])

	// fmt.Printf("fruitList: %v\n", fruitList)

	scores := make([]int, 4)

	scores = append(scores, 1)
	scores = append(scores, 2)
	scores = append(scores, 3)
	scores[len(scores)-1] = 2

	scores[0] = 2

	fmt.Printf("The scores are %v\n", scores)
	fmt.Printf("The scores are %v\n", sort.IntsAreSorted(scores))

	sort.Ints(scores)

	fmt.Printf("The scores are %v\n", scores)
	fmt.Printf("The scores are %v\n", sort.IntsAreSorted(scores))

	var intSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	index := 3

	intSlice = append(intSlice[:index], intSlice[index+1:]...)

	fmt.Printf("intSlice: %v\n", intSlice)

}
