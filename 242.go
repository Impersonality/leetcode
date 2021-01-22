package leetcode

// 桶排序，但是自己写的是80 30，因为用的是map，
// 看霜神题解，换成slice，100 20
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	byteS := []byte(s)
	byteT := []byte(t)
	m := make([]int, 26)
	for _, v := range byteS {
		m[v-'a']++
	}
	for _, v := range byteT {
		if m[v-'a'] < 1 {
			return false
		}
		m[v-'a']--
	}
	return true
}
