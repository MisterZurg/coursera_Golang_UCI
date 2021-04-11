/*
Write a program which prompts the user to enter a floating point number
and prints the integer which is a truncated version of the floating point number that was entered.
Truncation is the process of removing the digits to the right of the decimal place.
 */
package main

import "fmt"

func main() {
	var usersInput float64
	fmt.Println("Enter a floating point number:")
	fmt.Scan(&usersInput)
	fmt.Printf("Truncation of %.3f is %d",
		usersInput,
		int64(usersInput))
}
