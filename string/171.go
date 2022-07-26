package string

/*
个人思路：168的逆序题
答案：我想的是从低位向高位累加，实际高位向低位累加更好，每次循环都可以做点什么
*/

func titleToNumber(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		temp := int(s[i]-'A') + 1
		ans = ans*26 + temp
	}
	return ans
}
