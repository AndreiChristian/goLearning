package main

import (
	"aoc/utils"
	"fmt"
)

func main(){

	text := utils.Scan("input.txt")

	epsilon :=""
	gamma :=""

	len := len(text[0])	

	i:=0;

	for i< len{

		g, e := utils.Count(text, i)

		gamma = gamma + g

		epsilon = epsilon + e

		i ++
	}

	fmt.Println(gamma)
	fmt.Println(epsilon)

	gammaDecimal := utils.BinaryToDecimal(gamma)
	epsilonDecimal := utils.BinaryToDecimal(epsilon)

	fmt.Println(gammaDecimal)
	fmt.Println(epsilonDecimal)

	fmt.Println(gammaDecimal * epsilonDecimal)

}
