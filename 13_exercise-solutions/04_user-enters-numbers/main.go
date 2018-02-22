package main

import "fmt"

func main() {
	var lowNumber int
	fmt.Print("Please enter a low number: ")
	fmt.Scan(&lowNumber)

	var highNumber int
	fmt.Print("Please enter a higher number: ")
	fmt.Scan(&highNumber)

	result := highNumber / lowNumber
	fmt.Println("Large number divided by small number is:", result)
}
