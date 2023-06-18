package main

import "fmt"

func main() {

	myMap := make(map[string]string)

	myMap["js"] = "Javascript"
	myMap["ts"] = "Typescript"
	myMap["py"] = "Python"

	newMap := myMap

	myMap["js"] = "Java" //!

	newMap["go"] = "Golang"
	newMap["js"] = "Javascript"

	delete(myMap, "js")

	fmt.Println(myMap)
	fmt.Println(newMap)

	for k, v := range myMap {
		fmt.Printf("For key %v, value is %v.\n", k, v)
	}

}
