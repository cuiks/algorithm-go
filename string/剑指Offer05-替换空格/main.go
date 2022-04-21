package main

func replaceSpace(s string) string {
	result := ""
	for _, c := range s {
		switch c {
		case ' ':
			result += "%20"
		default:
			result += string(c)
		}
	}
	return result
}

func main() {

}
