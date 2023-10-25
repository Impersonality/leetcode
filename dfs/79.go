package dfs

/*
个人思路：要么dfs要么bfs，但是我完全忘记怎么实现了，直接看答案当复习吧
注意事项：visited数组是一个公有对象，需要重置它的值
*/

var dir = [][]int{
	{-1, 0},
	{1, 0},
	{0, 1},
	{0, -1},
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if existFunc(board, visited, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func inArea(board [][]byte, i, j int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	return true
}

func existFunc(board [][]byte, visited [][]bool, word string, index, i, j int) bool {
	if index == len(word)-1 {
		return board[i][j] == word[index]
	}
	if board[i][j] == word[index] {
		visited[i][j] = true
		for n := 0; n < 4; n++ {
			nx := i + dir[i][0]
			ny := i + dir[i][1]
			if inArea(board, nx, ny) && !visited[nx][ny] && existFunc(board, visited, word, index+1, nx, ny) {
				return true
			}
		}
		visited[i][j] = false
	}
	return false
}
