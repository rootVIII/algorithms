package main

import (
	"bytes"
	"fmt"
)

// rootVIII LCP in a single pass
// Leetcode 0 ms, faster than 100.00% of Go online submissions

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	var strbuf bytes.Buffer
	var lcp string = strs[0]

	for _, val := range strs[1:] {
		for pos, letter := range lcp {
			if pos >= len(val) || letter != rune(val[pos]) {
				break
			}
			strbuf.WriteRune(letter)
		}

		if strbuf.Len() > 0 {
			if strbuf.Len() < len(lcp) || strbuf.String() == val {
				lcp = strbuf.String()
			}
			strbuf.Reset()
		} else {
			return ""
		}
	}

	return lcp
}

func main() {
	fmt.Printf("%v\n", longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Printf("%v\n", longestCommonPrefix([]string{"cir", "car"}))
	fmt.Printf("%v\n", longestCommonPrefix([]string{"flower", "flow", "flight", "flom", "flam"}))
	fmt.Printf("%v\n", longestCommonPrefix([]string{"flower", "flower", "flower", "flower", "flower"}))
	fmt.Printf("%v\n", longestCommonPrefix([]string{"a", "a", "b"}))
}
