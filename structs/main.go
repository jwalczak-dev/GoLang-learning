package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{"Alex", "Anderson"}
	// bob := person{firstName: "Bob", lastName: "Sponge"}
	var joe person

	fmt.Println(alex)
	fmt.Println(joe)
	fmt.Printf("%+v", joe)

}
