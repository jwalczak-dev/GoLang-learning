package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type nameT string

func main() {
	// alex := person{"Alex", "Anderson"}
	// bob := person{firstName: "Bob", lastName: "Sponge"}
	var joe person

	// fmt.Println(alex)
	fmt.Println(joe)
	fmt.Printf("%+v\n", joe)

	joe.firstName = "Joe"
	joe.lastName = "Doe"
	fmt.Printf("%+v\n", joe)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.details()

	name := nameT("Jakub")
	name.updateN("Bob")
	fmt.Println(name)

}

func (n *nameT) updateN(newName string) {
	*n = nameT(newName)
}

func (pinterToPerson *person) updateName(newFirstName string) {
	(*pinterToPerson).firstName = newFirstName
}

func (p person) details() {
	fmt.Printf("%+v\n", p)
}
