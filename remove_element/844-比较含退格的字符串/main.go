package main

func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1
	sDel, tDel := 0, 0
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				sDel++
				i--
			} else if sDel > 0 {
				i--
				sDel--
			} else {
				break
			}
		}

		for j >= 0 {
			if t[j] == '#' {
				tDel++
				j--
			} else if tDel > 0 {
				j--
				tDel--
			} else {
				break
			}
		}

		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}

func main() {

}
