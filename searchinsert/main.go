package main

import (
	"fmt"
)

// rootVIII searchInsert (leetcode)

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)

}

func main() {
	fmt.Printf("\n%v\n", searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Printf("\n%v\n", searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Printf("\n%v\n", searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Printf("\n%v\n", searchInsert([]int{1, 3, 5, 6}, 0))
	fmt.Printf("\n%v\n", searchInsert([]int{1}, 0))
	fmt.Printf("\n%v\n", searchInsert([]int{1}, 1))
	fmt.Printf("\n%v\n", searchInsert([]int{1}, 2))
	fmt.Printf("\n%v\n", searchInsert([]int{1, 3, 5, 1}, 1))
}
