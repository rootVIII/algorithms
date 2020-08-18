package main

import (
	"fmt"

	"github.com/rootVIII/algorithms/mergesort"
)

func main() {

	// MergeSort
	fmt.Println("MergeSort")
	fmt.Printf("%v\n", mergesort.MergeSort([]uint8{8, 6, 3, 0, 4, 5, 3, 9, 2}))
	fmt.Println()

}
