package main

import "fmt"

type Person struct {
	FirstName     string
	LastName      string
	Age           int
	ContactDetail Contact
}

type Contact struct {
	Phone string
}

func main() {

	var roshan Person
	roshan.FirstName = "Roshan"
	roshan.LastName = "Poudel"
	roshan.Age = 28

	person := Person{
		FirstName: "Apple",
		LastName:  "Kaji",
		Age:       18,
	}

	person.ContactDetail = Contact{
		Phone: "+9779749742245",
	}

	person1 := new(Person)
	fmt.Println(roshan, person, person1)

}
