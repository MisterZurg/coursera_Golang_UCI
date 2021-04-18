/*
Write a Bubble Sort program in Go.
The program should prompt the user to type in a sequence of up to 10 integers.
The program should print the integers out on one line, in sorted order, from least to greatest.
Use your favorite search tool to find a description of how the bubble sort algorithm works.

As part of this program, you should write a function called BubbleSort()
which takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation
which swaps the position of two adjacent elements in the slice.
You should write a Swap() function which performs this operation.
Your Swap() function should take two arguments,
a slice of integers and an index value i which indicates a position in the slice.
The Swap() function should return nothing,
but it should swap the contents of the slice in position i with the contents in position i+1.
*/
package main

import (
	"fmt"
	"math/rand"
)

func Swap(ind1, ind2 int, slice *[]int) {
	var tmp int = (*slice)[ind1]
	(*slice)[ind1] = (*slice)[ind2]
	(*slice)[ind2] = tmp
}

func BubbleSort(sequence *[]int) {
	fmt.Println("Processing BubbleSort: O(n^2) complexity")
	// From "Алгоритмы сортировки в теории и на практике"
	var needIteration bool = true
	for needIteration {
		needIteration = false
		for i := 1; i < len(*sequence); i++ {
			if (*sequence)[i] < (*sequence)[i-1] {
				Swap(i, i-1, sequence)
				needIteration = true
			}
		}
	}
}

// Runs MergeSort algorithm on a slice single
func MergeSort(slice []int) []int {
	// Runtime: O(nlogn)
	// Memory: O(n)")
	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

// Merges left and right slice into newly created slice
func Merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

func QuickSort(slice []int) []int {
	// Алгоритмическая сложность: O(n log n)
	if len(slice) < 2 {
		return slice
	}

	left, right := 0, len(slice)-1

	pivot := rand.Int() % len(slice)

	slice[pivot], slice[right] = slice[right], slice[pivot]

	for i, _ := range slice {
		if slice[i] < slice[right] {
			slice[left], slice[i] = slice[i], slice[left]
			left++
		}
	}

	slice[left], slice[right] = slice[right], slice[left]

	QuickSort(slice[:left])
	QuickSort(slice[left+1:])

	return slice
}

func main() {
	var unsortedSlice = make([]int, 10)
	fmt.Println(unsortedSlice)
	for ind := range unsortedSlice {
		fmt.Scan(&unsortedSlice[ind])
	}

	fmt.Printf("Before BubbleSort %d\n", unsortedSlice)
	BubbleSort(&unsortedSlice)

	fmt.Printf("After BubbleSort %d\n", unsortedSlice)

	// sortedByMergeSort := MergeSort(unsortedSlice)
	// sortedByQuickSort := QuickSort(unsortedSlice)
	// fmt.Printf("After MergeSort %d\n",sortedByMergeSort)
	// fmt.Printf("After QuickSort %d\n",sortedByQuickSort)
}
