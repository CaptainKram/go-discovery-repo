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

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "mail",
			zipCode: 1000,
		},
	}

	var joe person
	joe.firstName = "Joe"
	joe.lastName = "Joeyson"
	joe.contactInfo = contactInfo{email: "mail", zipCode: 1000}

	joe.updateName("Jonny")

	alex.print()
	joe.print()

	mySlice := []string{"Hi", "There"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func updateSlice(s []string) {
	s[0] = "Hello"
}
