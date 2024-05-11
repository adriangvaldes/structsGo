package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email:   "jim@jim.com",
			zipCode: 94000,
		},
	}
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerPerson *person) updateName(newFirstName string) {
	(*pointerPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
