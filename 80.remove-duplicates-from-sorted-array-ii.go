package lc

/*
 * @lc app=leetcode.cn id=80 lang=golang
 * @lcpr version=21917
 *
 * [80] 删除有序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates2(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow := 2
	for fast := 2; fast < n; fast++ {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,1,2,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,1,1,1,1,2,3,3]\n
// @lcpr case=end

*/
