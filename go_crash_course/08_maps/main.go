package main

import "fmt"

func main() {
	//define map
	// emails := make(map[string]string)

	// //Asign kv
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// Declare map and add kv

	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	delete(emails, "Bob")
	fmt.Println(emails)
}
