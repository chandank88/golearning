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

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateFName(newFName string) {
	(*p).firstName = newFName
	//p.print()
}

func (p *person) updateLName(newLName string) {
	(*p).lastName = newLName
}

func main() {
	//alex := person{"Alex", "Anderson"}
	alex := person{firstName: "Alex", lastName: "Anderson", contactInfo: contactInfo{"a@b.com", 123345}}

	var alex1 person
	alex1.firstName = "Alex1"
	alex1.lastName = "Anderson1"
	alex1.contactInfo = contactInfo{"a1@b.com", 123}

	fmt.Println(alex)
	fmt.Println(alex1)

	jim := &alex
	jim.updateFName("Jim")
	alex.updateLName("Doe")
	alex.print()
}
