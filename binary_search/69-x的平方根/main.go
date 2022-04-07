package main

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	start, end := 1, x
	for start+1 < end {
		mid := (start + end) / 2
		value := mid * mid
		if value == x {
			return mid
		} else if value < x {
			start = mid
		} else {
			end = mid
		}
	}
	return start
}

func main() {

}
