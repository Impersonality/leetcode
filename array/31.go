package array

/*
思路：没思路，这题就是背题
答案：
题目有 3 个问题需要解决。如何找到下一个排列。不存在下一个排列的时候如何生成最小的排列。如何原地修改。先解决第一个问题，如何找到下一个排列。
下一个排列是找到一个大于当前排序的字典序，且变大的幅度最小。那么只能将较小的数与较大数做一次原地交换。并且较小数的下标要尽量靠右，较大数也要尽可能小。
原地交换以后，还需要将较大数右边的数按照升序重新排列。这样交换以后，才能生成下一个排列。以排列 [8,9,6,10,7,2] 为例：能找到的符合条件的
一对「较小数」与「较大数」的组合为 6 与 7，满足「较小数」尽量靠右，而「较大数」尽可能小。当完成交换后排列变为 [8,9,7,10,6,2]，此时我们可以重排「较小数」
右边的序列，序列变为 [8,9,7,2,6,10]。
第一步：在 nums[i] 中找到 i 使得 nums[i] < nums[i+1]，此时较小数为 nums[i]，并且 [i+1, n) 一定为下降区间。
第二步：如果找到了这样的 i ，则在下降区间 [i+1, n) 中从后往前找到第一个 j ，使得 nums[i] < nums[j] ，此时较大数为 nums[j]。
第三步，交换 nums[i] 和 nums[j]，此时区间 [i+1, n) 一定为降序区间。最后原地交换 [i+1, n) 区间内的元素，使其变为升序，无需对该区间进行排序。
如果第一步找不到符合条件的下标 i，说明当前序列已经是一个最大的排列。那么应该直接执行第三步，生成最小的排列。

上一个排列解法只需要改变两点：1.大于号换成小于号2.排降序
*/

func nextPermutation(nums []int) {
	i, j := 0, 0
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for j = len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		swap(&nums, i, j)
	}
	reverse31(&nums, i+1, len(nums)-1)
}

func lastPermutation(nums []int) {
	i, j := 0, 0
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for j = len(nums) - 1; j > i; j-- {
			if nums[j] < nums[i] {
				break
			}
		}
		swap(&nums, i, j)
	}
	reverse31(&nums, i+1, len(nums)-1)
}

func swap(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}

func reverse31(nums *[]int, i, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}
