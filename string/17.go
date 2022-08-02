package string

/*
个人思路：没有想到解法，看了提示DFS还是没解出来
答案：这种组合使用DFS，BFS都可以，本质都是递归
*/

var letterMap = []string{
	" ",    //0
	"",     //1
	"abc",  //2
	"def",  //3
	"ghi",  //4
	"jkl",  //5
	"mno",  //6
	"pqrs", //7
	"tuv",  //8
	"wxyz", //9
}

// BFS
func letterCombinations(digits string) []string {
	result := make([]string, 0)
	if len(digits) == 0 {
		return result
	}

	prefix := letterMap[digits[0]-'0']
	for _, v := range prefix {
		tail := letterCombinations(digits[1:])
		if len(tail) > 0 {
			for _, vv := range tail {
				result = append(result, string(v)+vv)
			}
		} else {
			result = append(result, string(v))
		}
	}
	return result
}

// DFS
func letterCombinations1(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	return letterFunc(digits, "")
}

func letterFunc(digits, prefix string) []string {
	result := make([]string, 0)
	if len(digits) == 0 {
		return []string{prefix}
	}
	for _, v := range letterMap[digits[0]-'0'] {
		result = append(result, letterFunc(digits[1:], prefix+string(v))...)
	}
	return result
}
