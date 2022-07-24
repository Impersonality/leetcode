package string

/*
个人思路：遍历，一个个比较，注意下corner case即可
答案：霜神是两个for遍历haystack和needle，不过写的很优雅，再次折服
*/

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if needle == haystack[i:i+len(needle)] {
			return i
		}
	}
	return -1
}
