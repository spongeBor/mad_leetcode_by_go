/*
 * @lc app=leetcode.cn id=2652 lang=golang
 * @lcpr version=30102
 *
 * [2652] 倍数求和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func f(n, m int) int {
	return (m + n/m*m) * (n / m) / 2
}

func sumOfMultiples(n int) int {
	return f(n, 3) + f(n, 5) + f(n, 7) - f(n, 3*5) - f(n, 3*7) - f(n, 5*7) + f(n, 3*5*7)
}

// @lc code=end

/*
// @lcpr case=start
// 7\n
// @lcpr case=end

// @lcpr case=start
// 10\n
// @lcpr case=end

// @lcpr case=start
// 9\n
// @lcpr case=end

*/

