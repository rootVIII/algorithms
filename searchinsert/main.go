package main

import (
	"fmt"
)

// rootVIII searchInsert (leetcode)
// 0 ms, faster than 100.00% of Go online submissions

func searchInsert(nums []int, target int) int {
	if (len(nums) == 1 && target == nums[0]) || target <= nums[0] {
		return 0
	}

	if target > nums[len(nums)-1] {
		return len(nums)
	}

	var i int = 0
	for ; ; i++ {

		if nums[i] < target && nums[i+1] >= target {
			i++
			break
		}
	}

	return i

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
