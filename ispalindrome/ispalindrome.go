package main

import (
	"fmt"
	"strings"
)

// IsPalindrome recursively checks if a string is a palindrome.
func IsPalindrome(phrase string) bool {
	if len(phrase) > 1 {
		if phrase[0] != phrase[len(phrase)-1] {
			return false
		}
		return IsPalindrome(phrase[1 : len(phrase)-1])
	}
	return true
}

func main() {

	var strArgs []string = []string{
		"Hello World",
		"racecar",
		"Was it a car or a cat I saw",
		"testing",
		"sit on a pan otis",
		"dude where's my car",
		"murder for a jar of red rum",
		"was it a rat I saw",
	}

	var edit string
	for _, arg := range strArgs {
		edit = strings.ToLower(strings.ReplaceAll(arg, " ", ""))
		fmt.Printf("%30s: %t\n", arg, IsPalindrome(edit))
	}
}
