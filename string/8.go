package string

/*
个人思路：这题没啥算法，考验的是实现代码能力，正常测试应该不会遇到。我通过率很低，用了12次，它的测试情况很多。
*/

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	result := 0
	sign := 1
	signEnable := true
	spaceEnable := true
	for _, v := range s {
		if v == ' ' && spaceEnable {
			continue
		}
		if v == '+' && signEnable {
			signEnable = false
			spaceEnable = false
			continue
		}
		if v == '-' && signEnable {
			signEnable = false
			spaceEnable = false
			sign = -1
			continue
		}

		signEnable = false
		spaceEnable = false
		c := v - '0'
		if c >= 0 && c <= 9 {
			newResult := result*10 + int(v-'0')
			if newResult < result {
				if sign == 1 {
					return (2 << 30) - 1
				}
				if sign == -1 {
					return -(2 << 30)
				}
			}
			result = newResult
		} else {
			break
		}
	}
	result = result * sign
	if result < -(2<<30) || result > (2<<30)-1 {
		if sign == -1 {
			result = -(2 << 30)
		} else {
			result = (2 << 30) - 1
		}
	}
	return result
}
