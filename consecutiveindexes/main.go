package main

import "fmt"

func consecutiveIndexes(items []string) []int {

	var matchesIndex []int
	var match string
	var likeMatches []string

	for i := 0; i < len(items); i++ {

		if items[i] != match {
			if len(likeMatches) > 1 {
				for j := i - len(likeMatches); j < i; j++ {
					matchesIndex = append(matchesIndex, j)
				}
			}

			likeMatches = nil
		}

		likeMatches = append(likeMatches, items[i])
		match = items[i]
	}

	return matchesIndex
}

func main() {

	// return indexes of any consecutive elements
	var ex1 = []string{
		"red", "red", "red", "red", "white", "blue",
		"blue", "yellow", "blue", "blue", "red",
		"purple", "purple", "purple", "purple", "red",
	}

	// 0 1 2 3 5 6 8 9 11 12 13 14
	fmt.Printf("%v\n", consecutiveIndexes(ex1))
}
