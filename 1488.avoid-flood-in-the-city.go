package lc

import "sort"

/*
 * @lc app=leetcode.cn id=1488 lang=golang
 * @lcpr version=21917
 *
 * [1488] 避免洪水泛滥
 */

// @lc code=start
func avoidFlood(rains []int) []int {
	// 获取输入数组的长度
	n := len(rains)
	// 创建一个与输入数组相同长度的整数切片，用于存储结果
	ans := make([]int, n)
	// 创建一个空切片 st，用于存储待抽水的湖泊的索引
	st := []int{}
	// 创建一个映射 mp，用于存储每个湖泊最后一次下雨的日子的索引
	mp := make(map[int]int)
	// 初始化结果切片，将其全部填充为 1，表示每天都可以抽水
	for i := 0; i < n; i++ {
		ans[i] = 1
	}
	// 遍历输入数组
	for i, rain := range rains {
		// 如果当前日子没有下雨（rain == 0）
		if rain == 0 {
			// 将当前日子的索引添加到 st 中，表示这个湖泊可以在未来某一天抽水
			st = append(st, i)
		} else {
			// 如果当前日子下雨（rain != 0），则将 ans[i] 设为 -1，表示无法抽水
			ans[i] = -1
			// 如果之前已经有雨水落在相同的湖泊上
			if day, ok := mp[rain]; ok {
				// 使用二分查找找到一个可以抽水的日子，将雨水抽走
				it := sort.SearchInts(st, day)
				// 如果没有找到可以抽水的日子，返回空切片，表示无法避免洪水
				if it == len(st) {
					return []int{}
				}
				// 将抽水的湖泊标记为当前湖泊（rain），并更新 st 切片
				ans[st[it]] = rain
				copy(st[it:len(st)-1], st[it+1:len(st)])
				st = st[:len(st)-1]
			}
			// 更新当前湖泊（rain）的最后一次下雨的日子的索引
			mp[rain] = i
		}
	}
	// 返回最终的结果切片
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,0,0,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,0,1,2]\n
// @lcpr case=end

*/
