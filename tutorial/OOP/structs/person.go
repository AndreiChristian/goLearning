package structs

type Person struct {
	firstName string
	lastName  string
	age       int
}

func New(firstName string, lastName string, age int) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}

}

func (p Person) Walk() string {

	return p.firstName + p.lastName + "likes walking \n"

}

func (p Person) GetFirstName() string {
	return p.firstName
}

func (p *Person) ChangeLastName(newLastName string) {

	p.lastName = newLastName

}
