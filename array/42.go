package array

/*
思路：只想到蛮力破解，没想到答案也是，这是道coding题
答案：
*/

// /////////////////////Row///////////////////////
func trapRow(height []int) int {
	sum := 0
	max := getMax(height)
	for i := 1; i <= max; i++ {
		var isStart bool
		tempSum := 0
		for j := 0; j < len(height); j++ {
			if isStart && height[j] < i {
				tempSum++
			}
			if height[j] >= i {
				sum += tempSum
				tempSum = 0
				isStart = true
			}
		}
	}
	return sum
}

func getMax(height []int) int {
	max := 0
	for _, item := range height {
		if item > max {
			max = item
		}
	}
	return max
}

// /////////////////////Column///////////////////////
func trapColumn(height []int) int {
	sum := 0
	for i := 0; i < len(height); i++ {
		leftMax, rightMax := 0, 0
		for left := 0; left < i; left++ {
			if height[left] > leftMax {
				leftMax = height[left]
			}
		}
		for right := i + 1; right < len(height); right++ {
			if height[right] > rightMax {
				rightMax = height[right]
			}
		}
		minMax := min(leftMax, rightMax)
		if minMax > height[i] {
			sum += minMax - height[i]
		}
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// /////////////////////DP///////////////////////
func trapDP(height []int) int {
	sum := 0
	leftBox, rightBox := make([]int, len(height)), make([]int, len(height))
	leftMax := 0
	for i := 0; i < len(height)-1; i++ {
		if height[i] > leftMax {
			leftMax = height[i]
		}
		leftBox[i+1] = leftMax
	}
	rightMax := 0
	for i := len(height) - 1; i > 0; i-- {
		if height[i] > rightMax {
			rightMax = height[i]
		}
		rightBox[i-1] = rightMax
	}
	for i := 0; i < len(height); i++ {
		minMax := min(leftBox[i], rightBox[i])
		if minMax > height[i] {
			sum += minMax - height[i]
		}
	}
	return sum
}

// /////////////////////TWO POINTS///////////////////////
func trapTwoPoints(height []int) int {
	sum := 0
	left, right := 0, len(height)-1
	maxLeft, maxRight := 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			}
			if height[left] < maxLeft {
				sum += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] > maxRight {
				maxRight = height[right]
			}
			if height[right] < maxRight {
				sum += maxRight - height[right]
			}
			right--
		}
	}
	return sum
}
