package lc

/*
 * @lc app=leetcode.cn id=55 lang=golang
 * @lcpr version=21917
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	var p int
	for i := range nums {
		if i > p {
			return false
		}
		p = maxInt(p, i+nums[i])
	}
	return true
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,1,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,1,0,4]\n
// @lcpr case=end

*/
