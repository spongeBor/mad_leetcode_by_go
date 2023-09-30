package lc

import "sort"

/*
 * @lc app=leetcode.cn id=2136 lang=golang
 * @lcpr version=21917
 *
 * [2136] 全部开花的最早一天
 */

// @lc code=start
func earliestFullBloom(plantTime, growTime []int) (ans int) {
	type pair struct{ p, g int }
	a := make([]pair, len(plantTime))
	for i, p := range plantTime {
		a[i] = pair{p, growTime[i]}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].g > a[j].g })
	days := 0
	for _, p := range a {
		days += p.p              // 累加播种天数
		ans = max(ans, days+p.g) // 再加上生长天数，就是这个种子的开花时间
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [1,4,3]\n[2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,2]\n[2,1,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n[1]\n
// @lcpr case=end

*/
