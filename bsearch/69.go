package linked

/*
容易想到二分法，二分法也好写。但在看了答案后，发现自己并没有对二分法掌握的很好，这里总结下二分法。
二分法的使用条件1.有序 2.有边界 3.可用下标取值。
二分法的思想使用l,r去逼近解，所以有些情况求得的是近似解。
for(l<r)或者for(l<=r)的过程中，如何退出循环呢，l<=mid<=r，所以在mid赋值给l,r时，需要增减一下数据，比如mid大了，可以r=mid-1，这就是不大于1的误差

这题更优解是牛顿迭代，思想是切线的解接近于曲线的解，不断用切线去逼近曲线的解。
构建f(x)-f(x0)=f'(x0)(x-x0)，然后用代码表示

该题进阶：精确到n位小数。上面其实已经分析了，二分法的误差就在mid赋值给l,r时，mid精准赋值是没有误差的，
但这样跳不出循环，所以mid+/-一个误差值再赋给l,r。如果是精确到0.001，r=mid-0.001/l=mid+0.001
*/

// //////////////////Binary Search////////////////////
func mySqrt(x int) int {
	left, right, mid := 0, x, (0+x)/2
	for left < right {
		if (mid*mid == x) || (mid*mid < x && (mid+1)*(mid+1) > x) {
			return mid
		}
		if mid*mid < x {
			left = mid + 1
		} else {
			right = mid
		}
		mid = (left + right) / 2
	}
	return mid
}

func mySqrt0(x int) int {
	l, r := 0, x
	for l < r {
		mid := (l + r + 1) / 2
		if mid*mid > x {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}

////////////////////Newton////////////////////
/*
f(x)-f(x0)=f'(x0)(x-x0)
x=x0-f(x0)/f'(x0)
*/
func mySqrt1(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}

// //////////////////精确到小数点n位////////////////////
func mySqrtFloat(x float64) float64 {
	l, r := 0.0, x
	for l <= r {
		mid := (l + r) / 2
		if x < mid*mid {
			r = mid - 1e-3
		} else {
			l = mid + 1e-3
		}
	}
	return r
}
