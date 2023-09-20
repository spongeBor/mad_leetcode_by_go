package lc

/*
 * @lc app=leetcode.cn id=LCP 06 lang=golang
 * @lcpr version=21914
 *
 * [LCP 06] 拿硬币
 */

// @lc code=start
func minCount(coins []int) int {
	sum := 0
	for _, i := range coins {
		sum += (i + 1) / 2
	}
	return sum
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,10]\n
// @lcpr case=end

*/
