package lc

/*
 * @lc app=leetcode.cn id=238 lang=golang
 * @lcpr version=21917
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}
	R := 1
	for i := length - 1; i >= 0; i-- {
		answer[i] = answer[i] * R
		R *= nums[i]
	}
	return answer
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [-1,1,0,-3,3]\n
// @lcpr case=end

*/
