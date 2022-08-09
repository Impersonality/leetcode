package string

/*
个人思路：没有想到解法，思路不清晰，没有找到突破口。问题可以转化为：子串在什么情况不能延续？什么时候计算子串长度？
答案：在多出一个)时，子串不延续。为了代码实现简洁，计算长度在每个)出现时，如果在子串结束时，那么s结束时也要判断。
	还有双指针解法，这个比较巧妙，我可能很难想出来，不过可以学习下这种思想，翻转
*/

func longestValidParentheses(s string) int {
	stack := []int{-1}
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 解法二 双指针
func longestValidParentheses1(s string) int {
	left, right, maxLength := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
}
