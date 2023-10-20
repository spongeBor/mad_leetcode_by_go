/*
 * @lc app=leetcode.cn id=2525 lang=golang
 * @lcpr version=30102
 *
 * [2525] 根据规则将箱子分类
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func categorizeBox(length, width, height, mass int) string {
	maxd := max(length, max(width, height))
	vol := length * width * height
	isBulky := maxd >= 10000 || vol >= 1e9
	isHeavy := mass >= 100
	if isBulky && isHeavy {
		return "Both"
	} else if isBulky {
		return "Bulky"
	} else if isHeavy {
		return "Heavy"
	} else {
		return "Neither"
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// 1000\n35\n700\n300\n
// @lcpr case=end

// @lcpr case=start
// 200\n50\n800\n50\n
// @lcpr case=end

*/

