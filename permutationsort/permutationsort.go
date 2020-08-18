package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Arrsort stores the incoming list to be sorted.
type Arrsort struct {
	incoming []int
}

// PermutationSort sorts a slice of integers.
func (a *Arrsort) PermutationSort() {
	for !a.isSorted() {
		a.shuffle()
	}
}

func (a Arrsort) isSorted() bool {
	for i := 0; i < len(a.incoming)-1; i++ {
		if a.incoming[i+1] < a.incoming[i] {
			return false
		}
	}
	return true
}

func (a *Arrsort) shuffle() {
	length := len(a.incoming)
	for i := 0; i < length; i++ {
		src := rand.NewSource(time.Now().UnixNano())
		newSrc := rand.New(src)
		r := newSrc.Intn(length - 1)
		tmp := a.incoming[i]
		a.incoming[i] = a.incoming[r]
		a.incoming[r] = tmp
	}
}

// GetSorted returns the sorted slice.
func (a Arrsort) GetSorted() []int {
	return a.incoming
}

func main() {

	// PermutationSort (really awful)
	fmt.Println("PermutationSort")
	var sorter = &Arrsort{incoming: []int{8, 1, 3, 4, 5, 2}}
	sorter.PermutationSort()
	fmt.Printf("%v\n", sorter.GetSorted())
	fmt.Println()
}
