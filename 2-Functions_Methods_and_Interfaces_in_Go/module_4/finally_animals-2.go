/*
Write a program which allows the user to create a set of animals
and to get information about those animals.
Each animal has a name and can be either a cow, bird, or snake.
With each command, the user can either create a new animal of one of the three types,
or the user can request information about an animal that he/she has already created.
Each animal has a unique name, defined by the user.
Note that the user can define animals of a chosen type,
but the types of animals are restricted to either cow, bird, or snake.
The following table contains the three types of animals and their associated data.

Animal 	| Food eaten | Locomotion method | Spoken sound
cow	   	| grass		 | walk				 | moo
bird	| worms		 | fly				 | peep
snake	| mice		 | slither			 | hsss

Your program should present the user with a prompt, “>”,
to indicate that the user can type a request.
Your program should accept one command at a time from the user,
print out a response, and print out a new prompt on a new line.
Your program should continue in this loop forever.
Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings.
The first string is “newanimal”.
The second string is an arbitrary string which will be the name of the new animal.
The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
Your program should process each newanimal command by creating the new animal
and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings.
The first string is “query”.
The second string is the name of the animal.
The third string is the name of the information requested
about the animal, either “eat”, “move”, or “speak”.
Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal.
Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(),
which take no arguments and return no values.
The Eat() method should print the animal’s food,
the Move() method should print the animal’s locomotion,
and the Speak() method should print the animal’s spoken sound.
Define three types Cow, Bird, and Snake.
For each of these three types, define methods Eat(), Move(), and Speak()
so that the types Cow, Bird, and Snake all satisfy the Animal interface.
When the user creates an animal, create an object of the appropriate type.
Your program should call the appropriate method when the user issues a query command.
*/
package main

import "fmt"

type Animal interface {
	getName() string
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (c Cow) getName() string {
	return c.name
}

func (c Cow) Eat() {
	fmt.Printf("%s eats grass\n", c.getName())
}

func (c Cow) Move() {
	fmt.Printf("%s walks\n", c.getName())
}

func (c Cow) Speak() {
	fmt.Printf("%s says moo\n", c.getName())
}

// ~~~~~~~~~~~~~~~~~~

type Bird struct {
	name string
}

func (b Bird) getName() string {
	return b.name
}

func (b Bird) Eat() {
	fmt.Printf("%s eats mice\n", b.getName())
}

func (b Bird) Move() {
	fmt.Printf("%s slithers\n", b.getName())
}

func (b Bird) Speak() {
	fmt.Printf("%s says peep\n", b.getName())
}

// ~~~~~~~~~~~~~~~~~~

type Snake struct {
	name string
}

func (s Snake) getName() string {
	return s.name
}

func (s Snake) Eat() {
	fmt.Printf("%s eats worms\n", s.getName())
}

func (s Snake) Move() {
	fmt.Printf("%s flies\n", s.getName())
}

func (s Snake) Speak() {
	fmt.Printf("%s says hsss\n", s.getName())
}

var Animals []Animal

func main() {
	// “newanimal” command or a “query”
	for {
		fmt.Println("What action you would like to do")
		fmt.Println("“newanimal” or “query”")
		var command, name, typeOrInformation string
		fmt.Scan(&command, &name, &typeOrInformation)
		switch command {
		case "newanimal":
			Animals = append(Animals, createAnimal(name, typeOrInformation))
		case "query":
			request(name, typeOrInformation)
		default:
			fmt.Println("Unexcpected Command Exception")
		}

		// fmt.Println(Animals)
	}
}

func createAnimal(name, tp string) Animal {
	switch tp {
	case "cow":
		fmt.Println("Created it!")
		return Cow{name: name}
	case "bird":
		fmt.Println("Created it!")
		return Bird{name: name}
	case "snake":
		fmt.Println("Created it!")
		return Snake{name: name}
	default:
		fmt.Println("Unexcpected Darwinism Exception")
		return nil
	}
}

func request(name, info string) {
	for _, creature := range Animals {
		if name == creature.getName() {
			fmt.Printf("Found %s\n", name)
			switch info {
			case "eat":
				creature.Eat()
			case "move":
				creature.Move()
			case "speak":
				creature.Speak()
			}
		}
	}
}
