package dfs

import "strings"

/*
才刷了一道dfs，所以这题比较轻松写出来了。
求排列组合的题都可以考虑dfs，dfs的写法是如果满足条件，返回递归到底了，判断是否满足条件。如果没递归到底，改变参数并继续递归
*/

func restoreIpAddresses(s string) []string {
	if s[0] == '0' {
		return restoreIpAddressesHelper("0", s[1:], 3)
	}
	res := make([]string, 0)
	if len(s) >= 3 && strings.Compare(s[:3], "255") <= 0 {
		res = append(res, restoreIpAddressesHelper(s[:3], s[3:], 3)...)
	}
	if len(s) >= 2 {
		res = append(res, restoreIpAddressesHelper(s[:2], s[2:], 3)...)
	}
	res = append(res, restoreIpAddressesHelper(s[:1], s[1:], 3)...)
	return res
}

func restoreIpAddressesHelper(prefix string, left string, num int) []string {
	if num == 0 && len(left) == 0 {
		return []string{prefix}
	}
	if num == 0 || len(left) == 0 {
		return nil
	}
	if left[0] == '0' {
		return restoreIpAddressesHelper(prefix+"."+"0", left[1:], num-1)
	}
	res := make([]string, 0)
	if len(left) >= 3 && strings.Compare(left[:3], "255") <= 0 {
		res = append(res, restoreIpAddressesHelper(prefix+"."+left[:3], left[3:], num-1)...)
	}
	if len(left) >= 2 {
		res = append(res, restoreIpAddressesHelper(prefix+"."+left[:2], left[2:], num-1)...)
	}
	res = append(res, restoreIpAddressesHelper(prefix+"."+left[:1], left[1:], num-1)...)
	return res
}
