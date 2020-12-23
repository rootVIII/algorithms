package main

import "fmt"

func lengthOfLastWord(s string) int {
	var wordlen int = 0
	var collecting bool = false

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' && !collecting {
			collecting = true
		}

		if collecting {
			if s[i] == ' ' {
				break
			} else {
				wordlen++
			}
		}

	}

	return wordlen
}

func main() {
	fmt.Println(lengthOfLastWord("Hellow World"))
	fmt.Println(lengthOfLastWord(" "))
	fmt.Println(lengthOfLastWord("a "))
}
