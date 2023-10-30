package lc

/*
 * @lc app=leetcode.cn id=274 lang=golang
 * @lcpr version=21917
 *
 * [274] H 指数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func hIndex(citations []int) int {
	// 答案最多只能到数组长度
	left, right := 0, len(citations)
	var mid int
	for left < right {
		// +1 防止死循环
		mid = (left + right + 1) >> 1
		cnt := 0
		for _, v := range citations {
			if v >= mid {
				cnt++
			}
		}
		if cnt >= mid {
			// 要找的答案在 [mid,right] 区间内
			left = mid
		} else {
			// 要找的答案在 [0,mid) 区间内
			right = mid - 1
		}
	}
	return left
}

// @lc code=end

/*
// @lcpr case=start
// [3,0,6,1,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,3,1]\n
// @lcpr case=end

*/
