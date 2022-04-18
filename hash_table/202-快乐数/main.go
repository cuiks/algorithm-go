package main

func isHappy(n int) bool {
	record := map[int]bool{}
	record[n] = true
	power := calPower(n)
	for power != 1 {
		if record[power] {
			return false
		}
		record[power] = true
		power = calPower(power)
	}
	return true
}

func calPower(n int) int {
	result := 0
	for n > 0 {
		v := n % 10
		result += v * v
		n /= 10
	}
	return result
}

func main() {

}
