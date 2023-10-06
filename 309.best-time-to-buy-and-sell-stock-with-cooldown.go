package lc

/*
 * @lc app=leetcode.cn id=309 lang=golang
 * @lcpr version=21917
 *
 * [309] 买卖股票的最佳时机含冷冻期
 */

// @lc code=start
func maxProfit6(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	f0, f1, f2 := -prices[0], 0, 0
	for i := 1; i < n; i++ {
		newf0 := max(f0, f2-prices[i])
		newf1 := f0 + prices[i]
		newf2 := max(f1, f2)
		f0, f1, f2 = newf0, newf1, newf2
	}
	return max(f1, f2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,0,2]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/
