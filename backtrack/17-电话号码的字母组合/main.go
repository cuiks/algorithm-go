package main

var keyMap map[byte]string = map[byte]string{
	'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
}

func letterCombinations(digits string) []string {
	result := make([]string, 0)
	if len(digits) == 0 {
		return result
	}
	backtrack(digits, 0, "", &result)
	return result
}

func backtrack(digits string, index int, combin string, result *[]string) {
	if len(digits) == len(combin) {
		*result = append(*result, combin)
		return
	}

	keys := keyMap[digits[index]]
	for i := 0; i < len(keys); i++ {
		backtrack(digits, index+1, combin+string(keys[i]), result)
	}
}

func main() {

}
