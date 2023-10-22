/*
 * @lc app=leetcode.cn id=1402 lang=golang
 * @lcpr version=30102
 *
 * [1402] 做菜顺序
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxSatisfaction(satisfaction []int) int {
	sort.Slice(satisfaction, func(i, j int) bool {
		return satisfaction[i] > satisfaction[j]
	})
	presum, ans := 0, 0
	for _, si := range satisfaction {
		if presum+si > 0 {
			presum += si
			ans += presum
		} else {
			break
		}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [-1,-8,0,5,-9]\n
// @lcpr case=end

// @lcpr case=start
// [4,3,2]\n
// @lcpr case=end

// @lcpr case=start
// [-1,-4,-5]\n
// @lcpr case=end

*/

