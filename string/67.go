package string

/*
个人思路：按照正常二进制加法思路，从末位加到首位
答案：两个char二进制相加，我的实现是if判断'1'，'0'，霜神是两数相加再加进位对2取余数和除数作为当前位和进位
	 数组a,b长度不相等，我会分别判断a>b和b>a的情况，霜神在函数开始if(a>b) a,b = b,a保证b>=a
*/

func addSingle(a, b uint8, carry bool) (bool, uint8) {
	if carry {
		if a == '1' && b == '1' {
			return true, '1'
		}
		if a == '1' || b == '1' {
			return true, '0'
		}
		return false, '1'
	}
	if a == '1' && b == '1' {
		return true, '0'
	}
	if a == '1' || b == '1' {
		return false, '1'
	}
	return false, '0'
}

func addBinary(a string, b string) string {
	lenA, lenB := len(a)-1, len(b)-1
	var carry bool
	var current uint8
	temp := make([]uint8, 0)
	for lenA >= 0 || lenB >= 0 {
		if lenA >= 0 && lenB >= 0 {
			carry, current = addSingle(a[lenA], b[lenB], carry)
		} else if lenA >= 0 {
			carry, current = addSingle(a[lenA], '0', carry)
		} else {
			carry, current = addSingle('0', b[lenB], carry)
		}
		temp = append(temp, current)
		lenA--
		lenB--
	}
	if carry {
		temp = append(temp, '1')
	}
	ans := ""
	for i := len(temp) - 1; i >= 0; i-- {
		ans += string(temp[i])
	}
	return ans
}
