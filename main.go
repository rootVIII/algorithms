package main

import "fmt"

func main() {
	for _, n := range Bubblesort1([]int{4, 3, 5, 8, 5, 0, 8, 18, 14}) {
		fmt.Printf("%d\n", n)
	}
	fmt.Println()

	var bsort = Bsort{incoming: []int{4, 3, 5, 8, 5, 0, 8, 18, 14}}
	bsort.Bubblesort2()
	for _, n := range bsort.incoming {
		fmt.Printf("%d\n", n)
	}
	fmt.Println()
}
