package leetcode

import "strconv"

/*
思路：想把所有数转成浮点，比如210->2.10，然后排序。不过浮点会比较占空间
答案：直接转成字符串比较，比较时需要处理：str(3) < str(30)，但是题目需要str(3) > str(30)
	处理："3" + "30" == "330" > "303"。排序用快排即可。
	corner case:给的全0数组，返回"0"
提交时超时：传统快排双指针lo, hi比较次数很多。改用霜神答案中的单指针，取hi为pivot，大于pivot依次放入数组前列，最后swap
			pivot和前列的最后一个
*/

func largestNumber(nums []int) string {
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}

	strBox := num2Str(nums)
	quickSort179(strBox)
	re := ""
	for _, v := range strBox {
		if re == "0" && v == "0" {
			continue
		}
		re += v
	}
	return re
}

func num2Str(nums []int) []string {
	strBox := make([]string, 0)
	for _, v := range nums {
		strBox = append(strBox, strconv.Itoa(v))
	}
	return strBox
}

func quickSort179(strBox []string) {
	if len(strBox) <= 1 {
		return
	}
	hight := len(strBox) - 1
	pivot := strBox[hight]

	i := -1
	for low := 0; low < hight; low++ {
		if strBox[low]+pivot > pivot+strBox[low] {
			i++
			strBox[i], strBox[low] = strBox[low], strBox[i]
		}
	}
	strBox[i+1], strBox[hight] = strBox[hight], strBox[i+1]
	quickSort179(strBox[:i+1])
	quickSort179(strBox[i+2:])
}
