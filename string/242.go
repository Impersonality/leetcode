package string

// 桶排序，但是自己写的是80 30，因为用的是map，
// 看霜神题解，换成slice，100 20

/*
个人思路：第二次写了，还是想的用map。应该养成思想：固定map->桶
答案：打表（桶排序）
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make([]int, 26)
	for _, v := range s {
		m[v-'a']++
	}
	for _, v := range t {
		if m[v-'a'] < 1 {
			return false
		}
		m[v-'a']--
	}
	return true
}
