package string

import "sort"

/*
个人思路：没解出来。做过判断两个字符串是否为anagram，使用map或者桶。这题就算给每个字符串一个map，如何判断两个map相等呢
答案：不应该被以往的题限制了思维，anagram的特点就是乱序，那么考虑先排序，排序后就迎刃而解了
*/

type bytes []byte

func (b bytes) Len() int {
	return len(b)
}

func (b bytes) Less(i int, j int) bool {
	return b[i] < b[j]
}

func (b bytes) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	count := make(map[string][]string)
	for _, s := range strs {
		byteS := bytes(s)
		sort.Sort(byteS)
		key := string(byteS)
		if _, ok := count[key]; !ok {
			count[key] = make([]string, 0)
		}
		count[key] = append(count[key], s)
	}
	for _, v := range count {
		res = append(res, v)
	}
	return res
}
