/*
 * @lc app=leetcode.cn id=1155 lang=golang
 * @lcpr version=30102
 *
 * [1155] 掷骰子等于目标和的方法数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numRollsToTarget(n int, k int, target int) int {
	mod := int(1e9 + 7)
	f := make([]int, target+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		for j := target; j >= 0; j-- {
			f[j] = 0
			for x := 1; x <= k; x++ {
				if j-x >= 0 {
					f[j] = (f[j] + f[j-x]) % mod
				}
			}
		}
	}
	return f[target]
}

// @lc code=end

/*
// @lcpr case=start
// 1\n6\n3\n
// @lcpr case=end

// @lcpr case=start
// 2\n6\n7\n
// @lcpr case=end

// @lcpr case=start
// 30\n30\n500\n
// @lcpr case=end

*/

