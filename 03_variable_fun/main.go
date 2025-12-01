package main

import "fmt"

func main() {
	// variables with explicit type
	var name string = "Suraj"
	var age int = 21

	// type inference
	salary := 55000.50
	isVerified := true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Salary:", salary)
	fmt.Println("Verified:", isVerified)
}
