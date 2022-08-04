package string

/*
个人思路：这题没啥算法，考验的是实现代码能力，我的思路是按照排列好的数组一个个添加到result
答案：霜神的要巧妙一些，用一个二维数组先存储char，拼接后输出
*/

func convert(s string, numRows int) string {
	result := ""
	round := numRows*2 - 2
	if numRows == 1 {
		round = 1
	}
	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += round {
			result += string(s[j])
			if (i%numRows) != 0 && (i%numRows) != numRows-1 && (j+round-2*i) < len(s) {
				result += string(s[j+round-2*i])
			}
		}
	}
	return result
}

///////////////////halfrost////////////////////////

func convert1(s string, numRows int) string {
	matrix, down, up := make([][]byte, numRows, numRows), 0, numRows-2
	for i := 0; i != len(s); {
		if down != numRows {
			matrix[down] = append(matrix[down], byte(s[i]))
			down++
			i++
		} else if up > 0 {
			matrix[up] = append(matrix[up], byte(s[i]))
			up--
			i++
		} else {
			up = numRows - 2
			down = 0
		}
	}
	solution := make([]byte, 0, len(s))
	for _, row := range matrix {
		for _, item := range row {
			solution = append(solution, item)
		}
	}
	return string(solution)
}
