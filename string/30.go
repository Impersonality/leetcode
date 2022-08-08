package string

/*
个人思路：先遍历s，然后依次与words比较，不过比较方法我想的是比较byte，效率太低
答案：循环比较应该使用map，这题其实考量的主要还是coding能力
*/

func findSubstring(s string, words []string) []int {
	count := make(map[string]int)
	for _, word := range words {
		count[word]++
	}
	tmpCount := copyMap(count)

	result := make([]int, 0)
	length := len(words[0])
	start := 0
	for i := 0; i+length <= len(s); i++ {
		if tmpCount[s[i:i+length]] > 0 {
			tmpCount[s[i:i+length]]--
			if checkWords(tmpCount) {
				result = append(result, start)
				continue
			}
			i = i + length - 1
		} else {
			start++
			i = start - 1
			tmpCount = copyMap(count)
		}
	}
	return result
}

func copyMap(source map[string]int) map[string]int {
	m := make(map[string]int)
	for k, v := range source {
		m[k] = v
	}
	return m
}

func checkWords(m map[string]int) bool {
	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}
