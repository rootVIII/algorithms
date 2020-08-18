package mergesort

/*
	rootVIII  |  30JAN2020
	merge sort 8-bit integers
*/

func merge(left []uint8, right []uint8) []uint8 {
	var combined []uint8
	leftIndex, rightIndex := 0, 0
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			combined = append(combined, left[leftIndex])
			leftIndex++
		} else {
			combined = append(combined, right[rightIndex])
			rightIndex++
		}
	}
	temp := append(left[leftIndex:], right[rightIndex:]...)
	combined = append(combined, temp...)
	return combined
}

// MergeSort sorts a slice of 8-bit unsigned integers.
func MergeSort(current []uint8) []uint8 {
	if len(current) < 2 {
		return current
	}
	m := int(len(current) / 2)
	return merge(MergeSort(current[0:m]), MergeSort(current[m:]))
}
