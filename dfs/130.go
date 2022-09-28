package dfs

/*
个人思路：将数组边界的O使用dfs转化为X，看错题了，题意是除了边界的O转化
答案：需要稍微转下弯，现在思维真是一点都不活跃了，可以先转化边界的O，然后剩下的就是满足题意的O了
*/

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	for i := range board {
		for j := range board[i] {
			if i == 0 || i == len(board)-1 || j == 0 || j == len(board[i])-1 {
				solveFunc(board, i, j)
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '*' {
				board[i][j] = 'O'
			}
		}
	}
}

func solveFunc(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}
	if board[i][j] == 'O' {
		board[i][j] = '*'
		solveFunc(board, i-1, j)
		solveFunc(board, i+1, j)
		solveFunc(board, i, j-1)
		solveFunc(board, i, j+1)
	}
}
