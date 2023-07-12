package snafu

import "math"

// Decimal          SNAFU
// 1              1
// 2              2
// 3             1=
// 4             1-
// 5             10
// 6             11
// 7             12
// 8             2=
// 9             2-
// 10             20
// 15            1=0
// 20            1-0
// 2022         1=11-2
// 12345        1-0---0
// 314159265  1121-1110-1=0

type Snafu struct {
	kv map[string]int
}

func New() Snafu {
	return Snafu{}
}

func (s *Snafu) SetMap() {

	newMap := make(map[string]int)

	newMap["2"] = 2
	newMap["1"] = 1
	newMap["0"] = 0
	newMap["-"] = -1
	newMap["="] = -2

	s.kv = newMap

}

func (s *Snafu) oneCharToInteger(key string) (int, error) {

	value, exists := s.kv[key]

	if exists {
		return value, nil
	}

	return -3, nil

}

func (s *Snafu) ToInteger(line []string) (int, error) {

	sum := 0

	for index, value := range line {

		parsedValue, err := s.oneCharToInteger(value)
		if err != nil {
			return 0, err
		}

		power := len(line) - index

		sum = sum + parsedValue*(int(math.Pow(5, float64(power-1))))

	}

	return sum, nil

}

// func (s *Snafu) oneCharFromInteger() string {}
