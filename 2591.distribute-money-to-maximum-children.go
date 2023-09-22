package lc

/*
 * @lc app=leetcode.cn id=2591 lang=golang
 * @lcpr version=21917
 *
 * [2591] 将钱分给最多的儿童
 */

// @lc code=start
func distMoney(money int, children int) int {
	if money < children {
		return -1
	}
	money -= children
	cnt := min(money/7, children)
	money -= cnt * 7
	children -= cnt
	if (children == 0 && money > 0) || (children == 1 && money == 3) {
		cnt -= 1
	}
	return cnt
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end

/*
// @lcpr case=start
// 20\n3\n
// @lcpr case=end

// @lcpr case=start
// 16\n2\n
// @lcpr case=end

*/
