package lc

import "sort"

/*
 * @lc app=leetcode.cn id=1333 lang=golang
 * @lcpr version=21917
 *
 * [1333] 餐厅过滤器
 */

// @lc code=start
func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	filtered := [][]int{}
	for _, r := range restaurants {
		if r[3] <= maxPrice && r[4] <= maxDistance && !(veganFriendly == 1 && r[2] == 0) {
			filtered = append(filtered, r)
		}
	}
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i][1] > filtered[j][1] || (filtered[i][1] == filtered[j][1] && filtered[i][0] > filtered[j][0])
	})
	res := []int{}
	for _, r := range filtered {
		res = append(res, r[0])
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[1,4,1,40,10],[2,8,0,50,5],[3,8,1,30,4],[4,10,0,10,3],[5,1,1,15,1]]\n1\n50\n10\n
// @lcpr case=end

// @lcpr case=start
// [[1,4,1,40,10],[2,8,0,50,5],[3,8,1,30,4],[4,10,0,10,3],[5,1,1,15,1]]\n0\n50\n10\n
// @lcpr case=end

// @lcpr case=start
// [[1,4,1,40,10],[2,8,0,50,5],[3,8,1,30,4],[4,10,0,10,3],[5,1,1,15,1]]\n0\n30\n3\n
// @lcpr case=end

*/
