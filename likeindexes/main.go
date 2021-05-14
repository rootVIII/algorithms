package main

import "fmt"

// rootVIII
// Return the index of each consecutive, matching string in the slice.
// There must be two or more consecutive matches.

func likeIndexes(collection []string) []int {

	var matchesIndex = make([]int, 0)
	var matchColor string
	var matchCount int

	for i, col := range collection {
		if col == matchColor {
			matchCount++
		} else {
			matchCount = 0
		}
		if matchCount > 0 && (i == len(collection)-1 || collection[i+1] != col) {
			for index := i - matchCount; index <= i; index++ {
				matchesIndex = append(matchesIndex, index)
			}
		}
		matchColor = col
	}
	return matchesIndex
}

func main() {
	// Expected: []int{2, 3, 4}
	fmt.Printf("%v\n", likeIndexes([]string{"one", "two", "three", "three", "three", "four", "five", "six", "seven", "eight"}))

	// Expected: []int{6, 7, 8, 9}
	fmt.Printf("%v\n", likeIndexes([]string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg", "ggg", "ggg", "ggg"}))

	// Expected: []int{}
	fmt.Printf("%v\n", likeIndexes([]string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg", "hhh", "iii", "jjj"}))

	// Expected: []int{0, 1, 5, 6}
	fmt.Printf("%v\n", likeIndexes([]string{"aaa", "aaa", "ccc", "ddd", "eee", "fff", "fff", "ggg", "hhh", "iii"}))

	// Expected: []int{0, 1, 2, 3, 7, 8}
	fmt.Printf("%v\n", likeIndexes([]string{"aaa", "aaa", "ccc", "ccc", "ddd", "eee", "fff", "eee", "eee", "iii"}))
}