package main

func minDeletionSize(strs []string) int {
	count := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] < strs[j-1][i] {
				count++
				break
			}
		}
	}
	return count
}

func main() {

}
