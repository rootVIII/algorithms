package main

import "fmt"

// https://leetcode.com/problems/single-number/

func singleNumber(nums []int) int {
	set := make(map[int]int, len(nums))
	for _, val := range nums {
		set[val]++
	}

	for key, val := range set {
		if val < 2 {
			return key
		}
	}
	return 0
}

func main() {
	fmt.Printf("%v\n", singleNumber([]int{2, 2, 1}))

	fmt.Printf("%v\n", singleNumber([]int{4, 1, 2, 1, 2}))

	fmt.Printf("%v\n", singleNumber([]int{1}))
}
