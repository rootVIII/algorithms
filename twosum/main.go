package main

import "fmt"

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)

	for i, n := range nums {
		diff := target - n
		contains := false
		for k := range m {
			if k != diff {
				continue
			}
			contains = true
		}
		if contains {

			return []int{m[diff], i}
		}
		m[n] = i
	}
	return []int{}
}

func main() {
	fmt.Printf("\n%v\n", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Printf("\n%v\n", twoSum([]int{3, 2, 4}, 6))
}
