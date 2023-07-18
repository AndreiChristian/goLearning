package main

import "fmt"

type ninjaStar struct {
	owner string
}

func (ninjaStar) attack() {
	fmt.Println("Throwing ninja star")
}

type ninjaSword struct {
	owner string
}

func (ninjaSword) attack() {

	fmt.Println("Using the ninja sword")

}

type ninjaWeapon interface {
	attack()
}

func Attack(weapon ninjaWeapon) {

	weapon.attack()

}

func main() {

	// stars := []ninjaStar{{owner: "Wallace"}, {owner: "Andrei"}}

	// swords := []ninjaSword{{owner: "Wallace"}, {owner: "Andrei"}}

	// for _, star := range stars {

	// 	star.attack()

	// }

	// for _, sword := range swords {
	// 	sword.attack()
	// }

	weapons := []ninjaWeapon{ninjaStar{owner: "Wallace"}, ninjaSword{owner: "Wallace"}}

	for _, weapon := range weapons {

		weapon.attack()

	}

	fmt.Print("The attack is finished \n")
}
