/*
 * @lc app=leetcode.cn id=2562 lang=golang
 * @lcpr version=21917
 *
 * [2562] 找出数组的串联值
 */

// @lc code=start
func findTheArrayConcVal(nums []int) int64 {
	ans := 0
	i, j := 0, len(nums)-1
	for i <= j {
		if i != j {
			val, _ := strconv.Atoi(strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]))
			ans += val
		} else {
			ans += nums[i]
		}
		i++
		j--
	}
	return int64(ans)
}

// @lc code=end

/*
// @lcpr case=start
// [7,52,2,4]\n
// @lcpr case=end

// @lcpr case=start
// [5,14,13,8,12]\n
// @lcpr case=end

*/

