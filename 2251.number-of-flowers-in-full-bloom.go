package lc

import "sort"

/*
 * @lc app=leetcode.cn id=2251 lang=golang
 * @lcpr version=21917
 *
 * [2251] 花期内花的数目
 */

// @lc code=start
func fullBloomFlowers(flowers [][]int, people []int) []int {
	n := len(flowers)
	starts := make([]int, n)
	ends := make([]int, n)
	for i, flower := range flowers {
		starts[i] = flower[0]
		ends[i] = flower[1]
	}
	sort.Ints(starts)
	sort.Ints(ends)
	m := len(people)
	ans := make([]int, m)
	for i, p := range people {
		x := sort.SearchInts(starts, p+1)
		y := sort.SearchInts(ends, p)
		ans[i] = x - y
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [[1,6],[3,7],[9,12],[4,13]]\n[2,3,7,11]\n
// @lcpr case=end

// @lcpr case=start
// [[1,10],[3,3]]\n[3,3,2]\n
// @lcpr case=end

*/
