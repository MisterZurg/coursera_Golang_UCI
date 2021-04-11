/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name,
in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name,
and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct which contains
the first and last names found in the file. Each struct created will be added to a slice,
and after all lines have been read from the file, your program will have a slice
containing one struct for each line in the file. After reading all lines from the file,
your program should iterate through your slice of structs and print
the first and last names found in each struct.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Human struct {
	fname string
	lname string
}

func main() {
	var fileName string
	var Humans []Human
	// prompt the user for the name of the text file.
	fmt.Scan(&fileName) // full_names
	file, err := os.Open(`1-Getting_Started_with_Go/module_4/res/` + fileName)
	// file, err := os.Open(`1-Getting_Started_with_Go/module_4/res/full_names`)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() { // возвращает true, пока файл не будет прочитан до конца
		line := sc.Text()
		fullname := strings.Split(line, " ")
		firstname := fullname[0]
		lastname := fullname[1]
		Humans = append(Humans, Human{
			fname: firstname,
			lname: lastname,
		})
	}
	fmt.Println(Humans)
}
