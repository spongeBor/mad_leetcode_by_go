package lc

/*
 * @lc app=leetcode.cn id=2582 lang=golang
 * @lcpr version=21917
 *
 * [2582] 递枕头
 */

// @lc code=start
func passThePillow(n int, time int) int {
	time %= (n - 1) * 2
	if time < n {
		return time + 1
	}
	return n - (time - (n - 1))
}

// @lc code=end

/*
// @lcpr case=start
// 4\n5\n
// @lcpr case=end

// @lcpr case=start
// 3\n2\n
// @lcpr case=end

*/
