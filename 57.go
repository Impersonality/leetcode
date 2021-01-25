package leetcode

// 57是56的加强版，但还是没写出来orz
// 思路和霜神的差不多：先添加newInterval之前的元素，然后newInterval怎么添加，怎么和后面的merge逻辑没捋清
// 写完后想了想，还是corner case捋不清，要把待解决问题拆开，这题，newInterval之前的算一类，非常干净
// newInterval要merge的算一类，不要merge前面和merge后面分开，很混乱，要尽可能用一个判断去分类，比如merge就算一个判断，
// 把newInterval放在最前，就不用考虑merge前面的元素

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	if len(intervals) == 0 {
		res = append(res, newInterval)
		return res
	}

	cur := 0
	for cur < len(intervals) && intervals[cur][1] < newInterval[0] {
		res = append(res, intervals[cur])
		cur++
	}

	for cur < len(intervals) && newInterval[1] >= intervals[cur][0] {
		newInterval = []int{min(intervals[cur][0], newInterval[0]), max(intervals[cur][1], newInterval[1])}
		cur++
	}
	res = append(res, newInterval)

	for cur < len(intervals) {
		res = append(res, intervals[cur])
		cur++
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
