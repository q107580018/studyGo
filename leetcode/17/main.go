package main

func main() {

}
func strStr(haystack string, needle string) int {
	n := len(needle)
	switch {
	case n == 0:
		return 0
	case n == len(haystack):
		if haystack == needle {
			return 0
		} else {
			return -1
		}
	case n > len(haystack):
		return -1

	case n < len(haystack):
		for i := 0; i < len(haystack)-len(needle)+1; i++ {
			flag := false
			if haystack[i] == needle[0] {
				flag = true
				for k, _ := range needle {
					if needle[k] != haystack[i+k] {
						flag = false
						break
					}
				}
			}
			if flag == true {
				return i
			}
		}
		return -1
	}
	return -1
}
