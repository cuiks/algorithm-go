package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	if len(s) < 4 || len(s) > 12 {
		return result
	}
	backtrack(s, 0, []string{}, &result)
	return result
}

func backtrack(s string, index int, list []string, result *[]string) {
	if len(list) == 4 {
		if index >= len(s) {
			*result = append(*result, strings.Join(list, "."))
		}
		return
	}
	for i := index; i < len(s); i++ {
		if v, err := strconv.Atoi(s[index : i+1]); err == nil && v >= 0 && v <= 255 {
			if i+1-index > 1 && s[index] == '0' {
				continue
			}
			list = append(list, s[index:i+1])
		} else {
			continue
		}
		backtrack(s, i+1, list, result)
		list = list[:len(list)-1]
	}
}

func main() {

}
