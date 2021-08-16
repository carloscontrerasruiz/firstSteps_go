package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(nFirstname string) {
	(*pointerToPerson).firstName = nFirstname
}

func main() {
	//alex := person{ "Alex", "Anderson"}
	/*alex := person{
	firstName: "Alex",
	lastName:  "Anderson"}*/

	/*var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"*/

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "email@email.com",
			zipCode: 1234,
		},
		contactInfo: contactInfo{
			email:   "email@email.com",
			zipCode: 1234,
		},
	}

	//fmt.Printf("%+v", alex)
	// Address to value *variable
	//Value to adress &variable
	alex.print()
	fmt.Println()
	//jimPointer := &alex
	alex.updateName("Jimbo")
	alex.print()
}
