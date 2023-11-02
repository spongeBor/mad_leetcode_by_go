/*
 * @lc app=leetcode.cn id=2103 lang=golang
 * @lcpr version=30104
 *
 * [2103] 环和杆
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
var POLE_NUM, COLOR_NUM = 10, 3

func countPoints(rings string) int {
	state := make([][]int, POLE_NUM)
	for i := 0; i < POLE_NUM; i++ {
		state[i] = make([]int, COLOR_NUM)
	}
	n := len(rings)
	for i := 0; i < n; i += 2 {
		color := rings[i]
		pole_index := rings[i+1] - '0'
		state[pole_index][getColorId(color)] = 1
	}
	res := 0
	for i := 0; i < POLE_NUM; i++ {
		flag := true
		for j := 0; j < COLOR_NUM; j++ {
			if state[i][j] == 0 {
				flag = false
				break
			}
		}
		if flag {
			res++
		}
	}
	return res
}

func getColorId(color byte) int {
	if color == 'R' {
		return 0
	} else if color == 'G' {
		return 1
	}
	return 2
}

// @lc code=end

/*
// @lcpr case=start
// "B0B6G0R6R0R6G9"\n
// @lcpr case=end

// @lcpr case=start
// "B0R0G0R9R0B0G0"\n
// @lcpr case=end

// @lcpr case=start
// "G4"\n
// @lcpr case=end

*/

