package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFName string) {
	(*pointerToPerson).firstName = newFName
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	alex.print()
	
	var laura person;
	laura.firstName = "Laura"
	laura.lastName = "Smith"

	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 96000,
		},
	}

	// this can be ommitted
	// b/c of the shortcut Go creates with 
	// pointers in recievers
	jimPointer := &jim
	jimPointer.updateName("jimmy")

	// with shortcut
	jim.updateName("Woazer")

	jim.print()
}

