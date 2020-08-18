package main

import (
	"fmt"

	"github.com/rootVIII/algorithms/bubblesort"
)

func main() {
	//Bubblesort1
	fmt.Println("BubbleSort1")
	for _, n := range bubblesort.BubbleSort1([]int{4, 3, 5, 8, 5, 0, 8, 18, 14}) {
		fmt.Printf("%d\n", n)
	}
	fmt.Println()

	// Bubblesort2
	fmt.Println("BubbleSort2")
	var bsort = &bubblesort.Bsort{Incoming: []int{4, 3, 5, 8, 5, 0, 8, 18, 14}}
	bsort.BubbleSort2()
	for _, n := range bsort.Incoming {
		fmt.Printf("%d\n", n)
	}
	fmt.Println()

}
