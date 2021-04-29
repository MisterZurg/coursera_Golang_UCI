/*
Write a program to sort an array of integers.
The program should partition the array into 4 parts,
each of which is sorted by a different goroutine.
Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers.
Each goroutine which sorts ¼ of the array should
*/

package main

import (
	"fmt"
	"log"
	"sort"
	"sync"
)

func main() {
	var amountOfNumbers, currentNumber int
	sequenceWithNumbers := make([]int, 0, 3)

	fmt.Println("Enter amount of number in sequence")
	_, err := fmt.Scan(&amountOfNumbers)
	if err != nil {
		log.Fatal(panic)
	}

	for i := 0; i < amountOfNumbers; i++ {
		fmt.Printf("Enter number №%d of %d\n", i+1, amountOfNumbers)
		_, err := fmt.Scan(&currentNumber)
		if err != nil {
			panic(err)
		}
		sequenceWithNumbers = append(sequenceWithNumbers, currentNumber)
	}

	fmt.Println("Your entered the folowing sequence: ", sequenceWithNumbers)

	var waitGroup sync.WaitGroup

	// Counting size for each of 4 subSlices
	var subSliceSize int = amountOfNumbers / 4
	subSlice1 := sequenceWithNumbers[0:subSliceSize]
	subSlice2 := sequenceWithNumbers[subSliceSize : 2*subSliceSize]
	subSlice3 := sequenceWithNumbers[2*subSliceSize : 3*subSliceSize]
	subSlice4 := sequenceWithNumbers[3*subSliceSize:]

	fmt.Printf("We split your sequence into \n%d\n%d\n%d\n%d\n", subSlice1, subSlice2, subSlice3, subSlice4)

	waitGroup.Add(4)
	go sliceSort(subSlice1)
	waitGroup.Done()

	go sliceSort(subSlice2)
	waitGroup.Done()

	go sliceSort(subSlice3)
	waitGroup.Done()

	go sliceSort(subSlice4)
	waitGroup.Done()

	waitGroup.Wait()

	finalSlice := mergeSubSlices(subSlice1, subSlice2, subSlice3, subSlice4)
	fmt.Println("My job here is done:", finalSlice)
}

func sliceSort(unsortSlice []int) []int {
	// fmt.Println("Before",unsortSlice)
	sort.Ints(unsortSlice)
	// fmt.Println("After",unsortSlice)
	return unsortSlice
}

func mergeSubSlices(subSlices ...[]int) []int {
	var outputSlice []int
	for _, subSlice := range subSlices {
		outputSlice = append(outputSlice, subSlice...)
	}
	// Sorting merged sorted parts
	return sliceSort(outputSlice)
}
