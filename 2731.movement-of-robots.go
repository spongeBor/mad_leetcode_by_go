/*
 * @lc app=leetcode.cn id=2731 lang=golang
 * @lcpr version=21917
 *
 * [2731] 移动机器人
 */

// @lc code=start
// 可以视碰撞为穿透
func sumDistance(nums []int, s string, d int) int {
	const mod int = 1e9 + 7
	n := len(nums)
	pos := make([]int, n)
	for i, ch := range s {
		if ch == 'L' {
			pos[i] = nums[i] - d
		} else {
			pos[i] = nums[i] + d
		}
	}
	sort.Ints(pos)
	res := 0
	for i := 1; i < n; i++ {
		res += (pos[i] - pos[i-1]) * i % mod * (n - i) % mod
		res %= mod
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [-2,0,2]\n"RLL"\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,0]\n"RL"\n2\n
// @lcpr case=end

*/

