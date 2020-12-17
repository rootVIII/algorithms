package main

import (
	"bytes"
	"fmt"
)

// rootVIII
// LCP in a single pass
// LEETCODE:
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Common Prefix.
// Memory Usage: 2.6 MB, less than 12.36% of Go online submissions for Longest Common Prefix.

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
		lastAllowedMatch := false
		for pos, letter := range lcp {
			if pos >= len(val) || letter != rune(val[pos]) {
				lastAllowedMatch = true
				continue
			}

			if !lastAllowedMatch {
				strbuf.WriteRune(letter)
			}
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
