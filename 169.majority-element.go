package lc

/*
 * @lc app=leetcode.cn id=169 lang=golang
 * @lcpr version=21917
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	major := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if major == num {
			count += 1
		} else {
			count -= 1
		}
	}

	return major

}

// @lc code=end

/*
// @lcpr case=start
// [3,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [2,2,1,1,1,2,2]\n
// @lcpr case=end

*/
