package leetcode

/*
思路：排序，然后对比下标和值(结果排序超时了。。。)
答案：也是排序... 是我快排写的不够好，用霜神的单指针模式吧，猜测是golang切片效率不如指针
*/

func hIndex(citations []int) int {
	if len(citations) < 1 {
		return 0
	}
	quickSort274(citations, 0, len(citations)-1)
	result := 0
	for k, v := range citations {
		if v >= k+1 {
			result++
		}
	}
	return result
}

func quickSort274(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}
	pivot := partition274(nums, lo, hi)
	quickSort274(nums, lo, pivot-1)
	quickSort274(nums, pivot+1, hi)
}

func partition274(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] > pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[hi], a[i+1] = a[i+1], a[hi]
	return i + 1
}
