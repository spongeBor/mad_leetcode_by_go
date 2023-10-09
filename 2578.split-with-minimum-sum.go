/*
 * @lc app=leetcode.cn id=2578 lang=golang
 * @lcpr version=21917
 *
 * [2578] 最小和分割
 */

// @lc code=start
func splitNum(num int) int {
	stnum := []byte(strconv.Itoa(num))
	sort.Slice(stnum, func(i, j int) bool {
		return stnum[i] < stnum[j]
	})
	num1, num2 := 0, 0
	for i := 0; i < len(stnum); i++ {
		if i%2 == 0 {
			num1 = num1*10 + int(stnum[i]-'0')
		} else {
			num2 = num2*10 + int(stnum[i]-'0')
		}
	}
	return num1 + num2
}

// @lc code=end

/*
// @lcpr case=start
// 4325\n
// @lcpr case=end

// @lcpr case=start
// 687\n
// @lcpr case=end

*/

