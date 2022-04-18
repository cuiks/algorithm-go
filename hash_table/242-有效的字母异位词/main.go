package main

func isAnagram(s string, t string) bool {
	maps := map[byte]int{}
	for i := 0; i < len(s); i++ {
		maps[s[i]]++
	}
	mapt := map[byte]int{}
	for i := 0; i < len(t); i++ {
		mapt[t[i]]++
	}
	if len(maps) != len(mapt) {
		return false
	}
	for k, v := range maps {
		if v != mapt[k] {
			return false
		}
	}
	return true
}

func main() {

}
