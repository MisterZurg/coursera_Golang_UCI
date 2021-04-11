/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and
address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map,
and then your program should print the JSON object.
*/
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string
	fmt.Println("Enter name: ")
	fmt.Scan(&name)
	fmt.Println("Enter address: ")
	fmt.Scan(&address)

	myMap := map[string]string{
		"name":    name,
		"address": address,
	}

	myPersonJson, err := json.Marshal(myMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(myPersonJson))
}
