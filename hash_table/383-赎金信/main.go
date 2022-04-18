package main

func canConstruct(ransomNote string, magazine string) bool {
	mMap := map[byte]int{}
	rMap := map[byte]int{}
	for i := 0; i < len(ransomNote); i++ {
		rMap[ransomNote[i]]++
	}
	for i := 0; i < len(magazine); i++ {
		mMap[magazine[i]]++
	}
	for k, v := range rMap {
		if mMap[k] < v {
			return false
		}
	}
	return true
}

func main() {

}
