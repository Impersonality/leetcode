package leetcode

import "math"

var N int
var result [][]string

func solveNQueens(n int) [][]string {
	N = n
	var result = make([][]string, 0)
	var row []int
	generateNQ(row, &result, 0, 0, 0)
	return result
}

func generateNQ(row []int, result *[][]string, column, diagonal1, diagonal2 int) {
	if len(row) == N {
		*result = append(*result, generateBoard(row))
		return
	}
	availablePosition := ((1 << N) - 1) & (^(column | diagonal1 | diagonal2))
	for availablePosition != 0 {
		position := availablePosition & (-availablePosition)
		t := math.Log2(float64(position))
		generateNQ(append(row, int(t)), result, column|position, (diagonal1|position)>>1, (diagonal2|position)<<1)
		availablePosition -= position
	}
}

func generateBoard(row []int) []string {
	var res []string
	for i := 0; i < N; i++ {
		temp := make([]byte, N)
		for j := 0; j < N; j++ {
			temp[j] = '.'
		}
		temp[row[i]] = 'Q'
		res = append(res, string(temp))
	}
	return res
}
