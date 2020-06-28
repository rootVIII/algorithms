package main

import "fmt"

func main() {
	//Bubblesort1
	fmt.Println("BubbleSort1")
	for _, n := range BubbleSort1([]int{4, 3, 5, 8, 5, 0, 8, 18, 14}) {
		fmt.Printf("%d\n", n)
	}
	fmt.Println()

	// Bubblesort2
	fmt.Println("BubbleSort2")
	var bsort = Bsort{incoming: []int{4, 3, 5, 8, 5, 0, 8, 18, 14}}
	bsort.BubbleSort2()
	for _, n := range bsort.incoming {
		fmt.Printf("%d\n", n)
	}
	fmt.Println()

	// MergeSort
	fmt.Println("MergeSort")
	fmt.Printf("%v\n", MergeSort([]uint8{8, 6, 3, 0, 4, 5, 3, 9, 2}))
	fmt.Println()

	// LCS
	fmt.Println("Longest Common Subsequence")
	var lcs = Lcs{}
	lcs.SortSequence([]uint8{8, 4, 9, 7, 8, 9, 3, 0, 4, 1})
	fmt.Printf("Shortest: %v\n", lcs.GetShortest())
	fmt.Printf("Longest: %v\n", lcs.GetLongest())
	fmt.Printf("All: %v\n", lcs.GetAll())
	fmt.Println()

	// InsertionSort
	fmt.Println("InsertionSort")
	unsortedInts := []int{8, 96, 34, 9, 18, 0, 8, 9, 51, 3, 21, 1}
	fmt.Printf("%v\n", InsertionSort(unsortedInts))
	fmt.Println()

	// InterpolationSearch
	fmt.Println("InterpolationSearch")
	bsort = Bsort{incoming: []int{4, 3, 5, 0, 8, 18, 14}}
	bsort.BubbleSort2()
	fmt.Printf("%d\n", InterpolationSearch(bsort.incoming, 8))
	fmt.Println()

	// PermutationSort (really awful)
	fmt.Println("PermutationSort")
	var sorter = &Arrsort{incoming: []int{8, 1, 3, 4, 5, 2}}
	sorter.PermutationSort()
	fmt.Printf("%v\n", sorter.GetSorted())
	fmt.Println()

}
