package main

import "fmt"

var name = "Firahs"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	//var name = "Firahs"

	var age = 31
	const isCool = true

	// Shorthand
	name := "Firahs"
	size := 1.3
	//email := "firahs@gmail.com"
	name, email := "Firahsz", "firahs@gmail.com"
	fmt.Println(name, age, isCool, size)
	fmt.Printf("%T\n", email)
}
