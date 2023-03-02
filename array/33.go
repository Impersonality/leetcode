package leetcode

/*
思路：数组分为两部分：大数组和小数组，若target>num[0]，从大数组中遍历；target<num[len]，从小数组中遍历。复杂度应该是O(n)，leetcode给的时间和空间复杂度都比较不错
答案：霜神答案：二分法，虽然并不能肯定target在low-mid还是mid-high，但是可以知道是不是在low-mid或者mid-high。还有个corner case：当mid=low或者mid=high时，low--/high++，避免mid死循环
*/

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r && target > nums[l] && (l == 0 || nums[l] >= nums[l-1]) {
		l++
	}
	if target == nums[l] {
		return l
	}
	for l < r && target < nums[r] && (r == len(nums)-1 || nums[r-1] <= nums[r]) {
		r--
	}
	if target == nums[r] {
		return r
	}
	return -1
}

//////////////////////////////////////////////////////////////////////////////////

func search33(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[low] { // 在数值大的一部分区间里
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] { // 在数值小的一部分区间里
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if nums[low] == nums[mid] {
				low++
			}
			if nums[high] == nums[mid] {
				high--
			}
		}
	}
	return -1
}
