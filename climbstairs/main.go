package main

import "fmt"

func climbStairs(n int) int {
	var items = []int{0, 1, 2}
	for i := 3; i < n+1; i++ {
		items = append(items, items[i-1]+items[i-2])
	}
	return items[n]
}

func main() {
	fmt.Println(climbStairs(45))
	fmt.Println(climbStairs(4))
}
