package string

/*
个人思路：没解出来。想尝试写isInteger, isDecimal函数拆解，不行
答案：用几个变量表示是否可以使用+-/./e
*/

func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}

	signEnable := true
	dotEnable := true
	eEnable := true
	for i := 0; i < len(s); i++ {
		if signEnable && (s[i] == '+' || s[i] == '-') && i < len(s)-1 {
			signEnable = false
			continue
		}
		if dotEnable && s[i] == '.' {
			if i-1 >= 0 && (s[i-1] >= '0' && s[i-1] <= '9') {
			} else if i+1 < len(s) && (s[i+1] >= '0' && s[i+1] <= '9') {
			} else {
				return false
			}
			dotEnable = false
			signEnable = false
			continue
		}
		if eEnable && (s[i] == 'e' || s[i] == 'E') {
			if i-1 >= 0 && ((s[i-1] >= '0' && s[i-1] <= '9') || s[i-1] == '.') && i+1 < len(s) {
				eEnable = false
				dotEnable = false
				signEnable = true
				continue
			}
		}
		if s[i] < '0' || s[i] > '9' {
			return false
		}
		signEnable = false
	}
	return true
}
