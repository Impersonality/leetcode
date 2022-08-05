package string

/*
个人思路：笨比写法，分别对1000/100/10/1层判断，再对4 9单独判断，再对 > 4 和 < 4判断
答案：贪心算法，对num递减，每次减掉最大数
*/

func intToRoman(num int) string {
	result := ""
	thousand := num / 1000
	num = num % 1000
	for i := 0; i < thousand; i++ {
		result += "M"
	}

	hundred := num / 100
	num = num % 100
	if hundred == 9 {
		result += "CM"
	} else if hundred == 4 {
		result += "CD"
	} else if hundred >= 5 {
		result += "D"
		for i := hundred - 5; i > 0; i-- {
			result += "C"
		}
	} else if hundred > 0 {
		for i := hundred; i > 0; i-- {
			result += "C"
		}
	}

	ten := num / 10
	num = num % 10
	if ten == 9 {
		result += "XC"
	} else if ten == 4 {
		result += "XL"
	} else if ten >= 5 {
		result += "L"
		for i := ten - 5; i > 0; i-- {
			result += "X"
		}
	} else if ten > 0 {
		for i := ten; i > 0; i-- {
			result += "X"
		}
	}

	if num == 9 {
		result += "IX"
	} else if num == 4 {
		result += "IV"
	} else if num >= 5 {
		result += "V"
		for i := num - 5; i > 0; i-- {
			result += "I"
		}
	} else if num > 0 {
		for i := num; i > 0; i-- {
			result += "I"
		}
	}

	return result
}

func intToRoman1(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res, i := "", 0
	for num != 0 {
		for values[i] > num {
			i++
		}
		num -= values[i]
		res += symbols[i]
	}
	return res
}
