package composition

type engine struct {
	hp int
}

type wheels struct {
	number int
}

type Car struct {
	engine engine
	wheels wheels
}

func NewCar() Car {

	return Car{
		engine: engine{100},
		wheels: wheels{4},
	}

}
