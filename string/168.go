package string

/*
个人思路：短除法，不过细节纠结了很久还是没想出来，我没想清楚26，52等能除尽的数如何处理
答案：我陷入了误区，在非除尽的数我也没算对，既然是26个连续的数，那么应该符合短除法，不会有特殊的数的处理
*/

func convertToTitle(columnNumber int) string {
	temp := make([]byte, 0)

	for columnNumber > 0 {
		temp = append(temp, 'A'+byte((columnNumber-1)%26))
		columnNumber = (columnNumber - 1) / 26
	}
	ans := ""
	for i := len(temp) - 1; i >= 0; i-- {
		ans += string(temp[i])
	}
	return ans
}
