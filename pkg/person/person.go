package person

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) Person {
	return Person{name: name, age: age}
}

func (person Person) GetName() string {
	return person.name
}

func (person Person) GetAge() int {
	return person.age
}
