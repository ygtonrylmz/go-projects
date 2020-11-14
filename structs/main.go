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
	onur := person{
		firstName: "Onur",
		lastName:  "Yılmaz",
		contactInfo: contactInfo{
			email:   "onur.y@thy.com",
			zipCode: 12345,
		},
	}
	onur.print()
	onur.updateName("Yiğit")
	onur.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
