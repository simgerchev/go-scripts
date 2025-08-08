package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello from my new Go application!")
	
	// Example: Simple greeting function
	name := "Developer"
	greeting := createGreeting(name)
	fmt.Println(greeting)
	
	// Example: Simple calculator
	a, b := 15, 3
	sum := add(a, b)
	product := multiply(a, b)
	
	fmt.Printf("%d + %d = %d\n", a, b, sum)
	fmt.Printf("%d * %d = %d\n", a, b, product)
}

func createGreeting(name string) string {
	return fmt.Sprintf("Welcome to Go programming, %s!", name)
}

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}
