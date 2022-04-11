package main

func totalFruit(fruits []int) int {
	totalMax := 0
	start, end := 0, 0
	total := 0
	totalKind := map[int]int{}
	for end < len(fruits) {
		totalKind[fruits[end]]++
		total++
		for len(totalKind) > 2 {
			totalKind[fruits[start]]--
			if totalKind[fruits[start]] == 0 {
				delete(totalKind, fruits[start])
			}
			start++
			total--
		}
		if total > totalMax {
			totalMax = total
		}
		end++
	}
	return totalMax
}

func main() {
}
