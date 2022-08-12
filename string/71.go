package string

import "strings"

/*
个人思路：没想出来
答案：其实没怎么想，工作繁琐，看到栈就豁然开朗了，coding难度不大，一两次就ac了
*/

func simplifyPath(path string) string {
	stack := make([]string, 0)
	arr := strings.Split(path, "/")
	for _, s := range arr {
		if s == "" || s == "/" || s == "." {
			continue
		}
		if s == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s)
		}
	}
	return "/" + strings.Join(stack, "/")
}
