package main

import "fmt"

func minWindow(s string, t string) string {
	result := s + s
	start, end := 0, 0
	sMap := map[byte]int{}
	tMap := map[byte]int{}
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}
	for end < len(s) {
		sMap[s[end]]++
		for check(sMap, tMap) {
			if end-start+1 < len(result) {
				result = s[start : end+1]
			}
			sMap[s[start]]--
			start++
		}
		end++
	}
	if result == s+s {
		return ""
	}
	return result
}

func check(sMap map[byte]int, tMap map[byte]int) bool {
	for key, value := range tMap {
		if sMap[key] < value {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(minWindow("a", "ABC"))
	//a := "sss"
	//fmt.Println(a[:3])

}
