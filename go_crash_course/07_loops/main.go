package main

import "fmt"

func main() {
	//long method
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}
	//short method

	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// FizzBuzz multiple 3 fizz, multiple 5 buzz, multiple 3 and 5 is fizzbuzz

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
