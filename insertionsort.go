package main

// InsertionSort sorts via insertion sort algorithm.
func InsertionSort(arr []int) []int {
	var i int
	var j int
	var key int

	for i = 1; i < int(len(arr)); i++ {
		key = arr[i]
		j = i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}
