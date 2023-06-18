package main

import "fmt"

func main() {

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	// fruitList[2] = "Peach"
	fruitList[3] = "Mango"

	// fruitList[5] = "Watermelon"

	fmt.Println(fruitList, len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "Carrot"}

	fmt.Println(vegList)

}
