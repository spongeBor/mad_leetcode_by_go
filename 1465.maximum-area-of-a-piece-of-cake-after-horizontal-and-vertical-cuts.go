/*
 * @lc app=leetcode.cn id=1465 lang=golang
 * @lcpr version=30102
 *
 * [1465] 切割后面积最大的蛋糕
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	return calMax(horizontalCuts, h) * calMax(verticalCuts, w) % 1000000007
}

func calMax(arr []int, boardr int) int {
	res, pre := 0, 0
	for _, i := range arr {
		res = max(i-pre, res)
		pre = i
	}
	return max(res, boardr-pre)
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

// @lc code=end

/*
// @lcpr case=start
// 5\n4\n[1,2,4]\n[1,3]\n
// @lcpr case=end

// @lcpr case=start
// 5\n4\n[3,1]\n[1]\n
// @lcpr case=end

// @lcpr case=start
// 5\n4\n[3]\n[3]\n
// @lcpr case=end

*/

