package main

// InsertionSort sorts via insertion sort algorithm.
func InsertionSort(arr []int) []int {
	var i int
	var j int
	var tmp int

	for i = 1; i < int(len(arr)); i++ {
		tmp = arr[i]
		j = i - 1
		for j >= 0 && arr[j] > tmp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = tmp
	}
	return arr
}
