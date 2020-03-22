package main

func main() {

}

func countCharacters(words []string, chars string) int {
	charsMap := make(map[rune]int)
	for _, v := range chars {
		charsMap[v] += 1
	}
	result := 0
	for _, word := range words {
		wordMap := make(map[rune]int)
		for _, char := range word {
			wordMap[char] += 1
		}
		flag := true
		for key, value := range wordMap {
			if charsMap[key] < value {
				flag = false
				break
			}
		}
		if flag {
			result += len(word)
		}

	}
	return result
}
