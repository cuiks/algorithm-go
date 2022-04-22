package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] && i+len(needle) <= len(haystack) && haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
func main() {
	fmt.Println(strStr("a", "a"))
}
