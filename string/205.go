package string

/*
个人思路：构建map对char一个个映射。对同构理解有偏差，第一次提交只从s映射了t，实际应该两侧都映射
答案：
*/

func f(s string, t string) bool {
	m := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = t[i]
		}
		if m[s[i]] != t[i] {
			return false
		}
	}
	return true
}

func isIsomorphic(s string, t string) bool {
	if !f(s, t) {
		return false
	}
	return f(t, s)
}
