package array

/*
思路：就想着双指针了，这题第一次见应该是做出来了，思维固化了
答案：构造map来存储每个数所需要的另一半
*/

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := m[nums[i]]; ok {
			return []int{val, i}
		}
		m[target-nums[i]] = i
	}
	return []int{}
}
