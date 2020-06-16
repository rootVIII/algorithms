package main

// bubblesort

func swap(first *int, second *int) {
	var temp int = *first
	*first = *second
	*second = temp
}

// Bubblesort1 sorts integers via pointer swap.
func Bubblesort1(incoming []int) []int {
	for i := 0; i < len(incoming)-1; i++ {
		for j := 0; j < len(incoming)-i-1; j++ {
			if incoming[j] > incoming[j+1] {
				swap(&incoming[j], &incoming[j+1])
			}
		}
	}
	return incoming
}
