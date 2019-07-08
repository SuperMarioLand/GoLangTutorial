package main

import (
	"fmt"
	"strconv"
)

// Person Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int

	//equivlent to firstName, lastName, city, gender string
	//age											 int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	// you can do this 1 million times and the person stays 26.wouldn't that be nice IRL
	// WTF google, spend all this time making goLang, why not make immortality real at this point?
	p.age++
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	}
	p.lastName = spouseLastName
}

func main() {
	//Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}

	// Alternative
	person2 := Person{"bob", "Jonson", "Boston", "m", 30}

	// fmt.Println(person1)
	// person1.age++
	// fmt.Println(person1)
	fmt.Println(person1.greet())
	fmt.Println(person1.greet())
	fmt.Println(person1.greet())

	person1.hasBirthday()
	person1.getMarried("Williams")
	fmt.Println(person1.greet())

	person2.getMarried("Thompson")
	person2.greet()
}
