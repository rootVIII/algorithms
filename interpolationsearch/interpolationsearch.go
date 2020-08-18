package main

import (
	"fmt"

	"github.com/rootVIII/algorithms/bubblesort"
)

// InterpolationSearch finds an integer.
func InterpolationSearch(sorted []int, target int) int {
	var low int = 0
	var high int = int(len(sorted)) - 1

	for low <= high && target >= sorted[low] && target <= sorted[high] {
		if low == high {
			return -1
		}
		position := low + (high-low)/(sorted[high]-sorted[low])*(target-sorted[low])
		if sorted[position] == target {
			return position
		}
		if sorted[position] < target {
			low = position + 1
		} else {
			low = position + 1
		}
	}
	return -1
}

// Bsort represents a slice of ints.
type Bsort struct {
	incoming []int
}

// BubbleSort2 sorts integers via pointer in method signature.
func (b *Bsort) BubbleSort2() {
	for i := 0; i < len(b.incoming)-1; i++ {
		for j := 0; j < len(b.incoming)-i-1; j++ {
			if b.incoming[j] > b.incoming[j+1] {
				temp := b.incoming[j]
				b.incoming[j] = b.incoming[j+1]
				b.incoming[j+1] = temp
			}
		}
	}
}

func main() {
	// InterpolationSearch
	fmt.Println("InterpolationSearch")
	bsort := bubblesort.Bsort{Incoming: []int{4, 3, 5, 0, 8, 18, 14}}
	bsort.BubbleSort2()
	// find position of number:
	fmt.Printf("%d\n", InterpolationSearch(bsort.Incoming, 8))
	fmt.Println()

}
