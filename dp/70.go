package dp

/*
思路：只想到了递归，递归来解斐波那契效率太低了。
答案：有时简单题最是难搞。一种解法是动态规划来模拟斐波那契，另一种是矩阵乘法
*/

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ////////////////////dp//////////////////////
func climbStairs(n int) int {
	i, j, k := 0, 1, 1
	for a := 0; a < n; a++ {
		k = i + j
		i = j
		j = k
	}
	return k
}

// ////////////////////matrix//////////////////////
type matrix [2][2]int

func mul(a, b matrix) matrix {
	return matrix{
		{a[0][0]*b[0][0] + a[0][1]*b[1][0], a[0][0]*b[1][0] + a[0][1]*b[1][1]},
		{a[1][0]*b[0][0] + a[1][1]*b[1][0], a[1][0]*b[1][0] + a[1][1]*b[1][1]},
	}
}

func climbStairsMatrix(n int) int {

	base := matrix{{1, 0}, {0, 0}}
	a := matrix{{1, 1}, {1, 0}}
	var res = a
	for i := 1; i < n; i++ {
		res = mul(res, a)
	}
	res = mul(res, base)
	return res[0][0]
}
