package array

// 思路：经典二分查找

func search704(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			r = mid - 1
		}
		if nums[mid] < target {
			l = mid + 1
		}
	}
	return -1
}
