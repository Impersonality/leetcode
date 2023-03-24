package array

/*
思路：双指针的思想容易想到，但是如果从头合并到nums1中，需要大量移动数组，没想出来方法。
答案：从尾开始填充即可，最后一个数组为空的处理：如果n==0，nums2已全部填入nums1，无需处理，m==0，将nums2剩下的填入nums1
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m + n; m > 0 && n > 0; i-- {
		if nums1[m-1] >= nums2[n-1] {
			nums1[i-1] = nums1[m-1]
			m--
		} else {
			nums1[i-1] = nums2[n-1]
			n--
		}
	}
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}
