package main

func main() {
}

func convertToTitle(n int) string {
	ans := []byte{}
	for ; n > 0; n = n / 26 {
		n--
		ans = append([]byte{'A' + byte(n%26)}, ans...)
	}
	return string(ans)
}
