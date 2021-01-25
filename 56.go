package leetcode

// 思路：两层for 每个元素与之前所有元素比较，是否需要merge
// 霜神答案，先对元素的左节点排序，然后遍历一次就可以merge了
// 以后遇到需要遍历多次时，考虑排序好是否能带来很大提升
// coding难度很大，首先是快排写不流畅，还是理解不到位
// 写完了直接超时

func QuickSort(arr [][]int) {
	if len(arr) < 2 {
		return
	}
	left := 0
	right := len(arr) - 1
	pivot := 0

	for left < right {
		for left < right && arr[pivot][0] <= arr[right][0] {
			right--
		}
		arr[pivot], arr[right] = arr[right], arr[pivot]
		pivot = right
		for left < right && arr[left][0] < arr[pivot][0] {
			left++
		}
		arr[pivot], arr[left] = arr[left], arr[pivot]
		pivot = left
	}
	QuickSort(arr[:pivot])
	QuickSort(arr[pivot+1:])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(arr [][]int) [][]int {
	if len(arr) < 2 {
		return arr
	}
	QuickSort(arr)
	res := make([][]int, 0)
	cur := 0
	res = append(res, arr[0])
	for i := 1; i < len(arr); i++ {
		if arr[i][0] > res[cur][1] {
			res = append(res, arr[i])
			cur++
		} else {
			res[cur][1] = max(res[cur][1], arr[i][1])
		}
	}
	return res
}
