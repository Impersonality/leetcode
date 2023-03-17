package leetcode

/*
思路：只想到蛮力破解，没想到答案也是，这是道coding题
答案：
*/

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	l, r, u, d := 0, len(matrix[0])-1, 0, len(matrix)-1
	for direction := 0; l <= r && u <= d; direction++ {
		switch direction % 4 {
		case 0:
			for i := l; i <= r; i++ {
				res = append(res, matrix[u][i])
			}
			u++
		case 1:
			for i := u; i <= d; i++ {
				res = append(res, matrix[i][r])
			}
			r--
		case 2:
			for i := r; i >= l; i-- {
				res = append(res, matrix[d][i])
			}
			d--
		case 3:
			for i := d; i >= u; i-- {
				res = append(res, matrix[i][l])
			}
			l++
		}
	}
	return res
}
