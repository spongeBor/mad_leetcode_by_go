package lc

/*
 * @lc app=leetcode.cn id=137 lang=golang
 * @lcpr version=30102
 *
 * [137] 只出现一次的数字 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func singleNumber(nums []int) int {
	freq := map[int]int{}
	for _, v := range nums {
		freq[v]++
	}
	for num, occ := range freq {
		if occ == 1 {
			return num
		}
	}
	return 0 // 不会发生，数据保证有一个元素仅出现一次
}

// @lc code=end

/*
// @lcpr case=start
// [2,2,3,2]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,0,1,0,1,99]\n
// @lcpr case=end

*/
