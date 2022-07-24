package string

/*
个人思路：动态规划，将字符串分为最后一个(最前一个)字符和其他字符，然后递归，测试后发现行不通，不能将其他字符看成一个整体
答案：思想太僵化了，按照人正常的思路然后用代码实现即可，从左至右，根据相邻两个char判断当前正负值，最后一个char正值
*/

var char2int = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return char2int[s[0]]
	}
	sum := 0
	for i := 0; i < len(s)-1; i++ {
		left := char2int[s[i]]
		right := char2int[s[i+1]]
		if left < right {
			sum -= left
		} else {
			sum += left
		}
	}
	sum += char2int[s[len(s)-1]]
	return sum
}
