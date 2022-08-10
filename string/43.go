package string

/*
个人思路：想遍历字符串，一位位的转换char->int 然后乘起来。这样不行，和直接转为str->int没区别
答案：创建一个数组长度为 len(num1) + len(num2) 的数组用于存储乘积。
     corner case:如果是从低位开始相乘，>10进位可行，如果是高位开始，建议先不进位，因为低位的进位可能会再次影响高位，等乘完了，遍历进位
                 然后要清除开始的0
*/

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	product := make([]int, len(num1)+len(num2))
	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			temp := int((num1[i] - '0') * (num2[j] - '0'))
			product[i+j+1] += temp
		}
	}

	for i := len(product) - 1; i > 0; i-- {
		product[i-1] += product[i] / 10
		product[i] = product[i] % 10
	}
	if product[0] == 0 {
		product = product[1:]
	}

	result := make([]byte, len(product))
	for i := 0; i < len(product); i++ {
		result[i] = '0' + byte(product[i])
	}
	return string(result)
}
