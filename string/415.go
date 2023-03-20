package string

import "strconv"

/*
个人思路：这题应该做过，但这次没写出来。思路容易想到，两两byte相加进位。但是coding能力太弱，一想到很多corner case，直接放弃了
答案：
*/
func addStrings(num1 string, num2 string) string {
	res := ""
	var carry uint8 = 0
	i, j := len(num1)-1, len(num2)-1
	for i >= 0 || j >= 0 {
		var intI, intJ uint8
		if i >= 0 {
			intI = num1[i] - '0'
		}
		if j >= 0 {
			intJ = num2[j] - '0'
		}
		i--
		j--
		tmp := intI + intJ + carry
		carry = tmp / 10
		remainder := tmp % 10
		res = strconv.Itoa(int(remainder)) + res
	}
	if carry > 0 {
		res = "1" + res
	}
	return res
}
