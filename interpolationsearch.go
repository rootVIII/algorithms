package main

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
