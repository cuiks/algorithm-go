package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	resultMap := map[string][]string{}
	for i := 0; i < len(strs); i++ {
		newStr := reorderStr(strs[i])
		resultMap[newStr] = append(resultMap[newStr], strs[i])
	}
	result := make([][]string, 0)
	for _, v := range resultMap {
		result = append(result, v)
	}
	return result
}

func reorderStr(str string) string {
	s := []byte(str)
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return string(s)
}

func main() {

}
