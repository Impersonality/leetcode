package leetcode

/*
思路：无，想到了先排序，然后排序保存原有下标（思路混乱。。）
答案：滑动窗口，遍历数组，每次滑动比较t次
	 桶排序，也是滑动窗口，每次滑动比较次数减少，只有1
corner case:一个桶的大小为t+1，负数桶的key需要-1
			 k == 0 时为false
*/

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) <= 1 || k == 0 {
		return false
	}
	bucket := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		key := calKey(nums[i], t)
		if _, ok := bucket[key]; ok {
			return true
		}

		if value, ok := bucket[key+1]; ok && abs(nums[i]-value) <= t {
			return true
		}
		if value, ok := bucket[key-1]; ok && abs(nums[i]-value) <= t {
			return true
		}

		if len(bucket) >= k {
			delete(bucket, calKey(nums[i-k], t))
		}
		bucket[key] = nums[i]
	}
	return false
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func calKey(val, t int) int {
	if val >= 0 {
		return val / (t + 1)
	} else {
		return (val / (t + 1)) - 1
	}
}
