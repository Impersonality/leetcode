package array

/*
思路：有想到第二种方法，双指针一个个查找，直到中位数，不过复杂度O(m+n).不过没写出来，这题corner case很复杂。
答案：答案的两种O(log(m+n))的算法理解不难，但实现起来也是非常复杂，这题算是先抄过一遍，之后再来吧
*/

// ////////////////////遍历查找 O(m+n)//////////////////////
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	length := m + n
	left, right := -1, -1
	aStart, bStart := 0, 0
	for i := 0; i <= length/2; i++ {
		left = right
		if aStart < m && (bStart >= n || nums1[aStart] <= nums2[bStart]) {
			right = nums1[aStart]
			aStart++
		} else {
			right = nums2[bStart]
			bStart++
		}
	}
	if length%2 == 1 {
		return float64(right)
	} else {
		return (float64(left) + float64(right)) / 2
	}
}

// ////////////////////getKth O(long(m+n))//////////////////////
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	left := (m + n + 1) / 2
	right := (m + n + 2) / 2
	leftRes := getKth(nums1, 0, m-1, nums2, 0, n-1, left)
	rightRes := getKth(nums1, 0, m-1, nums2, 0, n-1, right)
	return float64(leftRes+rightRes) / 2
}

func getKth(nums1 []int, start1, end1 int, nums2 []int, start2, end2 int, k int) int {
	length1, length2 := end1-start1+1, end2-start2+1
	if length1 == 0 {
		return nums2[start2+k-1]
	}
	if length2 == 0 {
		return nums1[start1+k-1]
	}
	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}
	i := start1 + min(length1, k/2) - 1
	j := start2 + min(length2, k/2) - 1
	if nums1[i] < nums2[j] {
		return getKth(nums1, i+1, end1, nums2, start2, end2, k-(i-start1+1))
	} else {
		return getKth(nums1, start1, end1, nums2, j+1, end2, k-(j-start2+1))
	}
}

// ////////////////////Binary Search O(long(m+n))//////////////////////
func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return findMedianSortedArrays3(nums2, nums1)
	}
	iMin, iMax := 0, m
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := (m+n+1)/2 - i
		if j != 0 && i != m && nums1[i] < nums2[j-1] {
			iMin = i + 1
		} else if i != 0 && j != n && nums2[j] < nums1[i-1] {
			iMax = i - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = Max(nums1[i-1], nums2[j-1])
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = min(nums1[i], nums2[j])
			}
			return float64(maxLeft+minRight) / 2
		}
	}
	return 0
}
