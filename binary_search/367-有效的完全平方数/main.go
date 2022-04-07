package main

func isPerfectSquare(num int) bool {
	start, end := 0, num
	for start+1 < end {
		mid := (start + end) / 2
		value := mid * mid
		if value == num {
			return true
		} else if value < num {
			start = mid
		} else {
			end = mid
		}
	}

	if start*start == num || end*end == num {
		return true
	}
	return false
}

func main() {
}
