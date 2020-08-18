package bubblesort

// bubblesort

func swap(first *int, second *int) {
	var temp int = *first
	*first = *second
	*second = temp
}

// BubbleSort1 sorts integers via pointer swap.
func BubbleSort1(incoming []int) []int {
	for i := 0; i < len(incoming)-1; i++ {
		for j := 0; j < len(incoming)-i-1; j++ {
			if incoming[j] > incoming[j+1] {
				swap(&incoming[j], &incoming[j+1])
			}
		}
	}
	return incoming
}

// Bsort represents a slice of ints.
type Bsort struct {
	Incoming []int
}

// BubbleSort2 sorts integers via pointer in method signature.
func (b *Bsort) BubbleSort2() {
	for i := 0; i < len(b.Incoming)-1; i++ {
		for j := 0; j < len(b.Incoming)-i-1; j++ {
			if b.Incoming[j] > b.Incoming[j+1] {
				temp := b.Incoming[j]
				b.Incoming[j] = b.Incoming[j+1]
				b.Incoming[j+1] = temp
			}
		}
	}
}
