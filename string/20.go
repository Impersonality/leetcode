package string

/*
个人思路：栈，这题没啥意思，leetcode记录上我还做过。不过go的容器不完善，不导包没有队列。我临时没有想到好的替代栈的方案。
答案：用一个数组暂存，手动实现栈的push和pop
*/

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, v := range s {
		if v == '{' || v == '[' || v == '(' {
			stack = append(stack, v)
		} else if len(stack) > 0 && ((v == '}' && stack[len(stack)-1] == '{') ||
			(v == ']' && stack[len(stack)-1] == '[') ||
			(v == ')' && stack[len(stack)-1] == '(')) {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
