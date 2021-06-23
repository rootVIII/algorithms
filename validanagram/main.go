package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var table [26]int
	i := 0
	for i = 0; i < len(s); i++ {
		table[rune(s[i])-'a']++
	}
	for i = 0; i < len(t); i++ {
		table[rune(t[i])-'a']--
		if table[rune(t[i])-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {
	x := "theta"
	y := "heatt"
	fmt.Println(isAnagram(x, y))

	x = "test"
	y = "ests"
	fmt.Println(isAnagram(x, y))

	x = "helloworld"
	y = "ellhowrold"
	fmt.Println(isAnagram(x, y))
}
