package utils

import (
	"log"
	"math"
	"strconv"
)

func BinaryToDecimal(binary string) int {

	b, err := strconv.Atoi(binary)

	if err!=nil{
		log.Fatal(err)
	}

	decimal :=0
	index :=0

	for b >0 {

		lastDigit := b%10;

		decimal = decimal + lastDigit* int(math.Pow(2,float64(index)))

		b = b/10
		index ++

	}

	return decimal

}
