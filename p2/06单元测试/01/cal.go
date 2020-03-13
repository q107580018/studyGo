package cal

func AddUpper(a int) int {
	sum := 0
	for ; a > 0; a-- {
		sum += a
	}
	return sum
}

func GetSub(a, b int) int {
	return a - b
}
