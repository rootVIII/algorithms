package main

import "fmt"

func main() {
	//Bubblesort1
	fmt.Println("Bubblesort1")
	for _, n := range BubbleSort1([]int{4, 3, 5, 8, 5, 0, 8, 18, 14}) {
		fmt.Printf("%d\n", n)
	}
	fmt.Println()

	// Bubblesort2
	fmt.Println("Bubblesort2")
	var bsort = Bsort{incoming: []int{4, 3, 5, 8, 5, 0, 8, 18, 14}}
	bsort.BubbleSort2()
	for _, n := range bsort.incoming {
		fmt.Printf("%d\n", n)
	}
	fmt.Println()

	// MergeSort
	fmt.Println("Mergesort")
	unsorted := []uint8{8, 6, 3, 0, 4, 5, 3, 9, 2}
	fmt.Printf("%v\n", MergeSort(unsorted))
	fmt.Println()

	// LCS
	fmt.Println("Longest Common Subsequence")
	unsorted = []uint8{8, 4, 9, 7, 8, 9, 3, 0, 4, 1}
	var lcs = Lcs{}
	lcs.SortSequence(unsorted)
	fmt.Printf("Shortest: %v\n", lcs.GetShortest())
	fmt.Printf("Longest: %v\n", lcs.GetLongest())
	fmt.Printf("All: %v\n", lcs.GetAll())
	fmt.Println()
}
