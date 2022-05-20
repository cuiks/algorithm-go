package main

func lemonadeChange(bills []int) bool {
	countFive := 0
	countTen := 0
	for _, v := range bills {
		if v == 5 {
			countFive++
		} else if v == 10 {
			countTen++
			countFive--
		} else {
			if countTen > 0 {
				countTen--
				countFive--
			} else {
				countFive -= 3
			}
		}
		if countTen < 0 || countFive < 0 {
			return false
		}
	}
	return true
}

func main() {

}
