/*
 * @lc app=leetcode.cn id=260 lang=golang
 * @lcpr version=30102
 *
 * [260] 只出现一次的数字 III
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func singleNumber(nums []int) (ans []int) {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}
	for num, occ := range freq {
		if occ == 1 {
			ans = append(ans, num)
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,1,3,2,5]\n
// @lcpr case=end

// @lcpr case=start
// [-1,0]\n
// @lcpr case=end

// @lcpr case=start
// [0,1]\n
// @lcpr case=end

*/

