package leetcode

// 抄官方答案，说是排序，不是常规排序，桶排序
// 我理解桶就是利用byte或者枚举来装输入
// 凡是字符串切分成byte都可以考虑
func sortString(s string) string {
	bucket := [26]int{}
	for _, v := range s {
		bucket[v-'a']++
	}

	re := make([]byte, 0, len(s))
	for len(re) < len(s) {
		for i := 'a'; i <= 'z'; i++ {
			if bucket[i-'a'] > 0 {
				re = append(re, byte(i))
				bucket[i-'a']--
			}
		}
		for i := 'z'; i >= 'a'; i-- {
			if bucket[i-'a'] > 0 {
				re = append(re, byte(i))
				bucket[i-'a']--
			}
		}
	}
	return string(re)
}
