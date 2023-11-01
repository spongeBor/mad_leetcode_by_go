/*
 * @lc app=leetcode.cn id=2127 lang=golang
 * @lcpr version=30103
 *
 * [2127] 参加会议的最多员工数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maximumInvitations(favorite []int) int {
	n := len(favorite)
	// 统计入度，便于进行拓扑排序
	indeg := make([]int, n)
	for _, x := range favorite {
		indeg[x]++
	}

	used := make([]bool, n)
	f := make([]int, n)
	for i := 0; i < n; i++ {
		f[i] = 1
	}

	q := []int{}
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		u := q[0]
		used[u] = true
		q = q[1:]
		v := favorite[u]
		// 状态转移
		f[v] = max(f[v], f[u]+1)
		indeg[v]--
		if indeg[v] == 0 {
			q = append(q, v)
		}
	}
	// ring 表示最大的环的大小
	// total 表示所有环大小为 2 的「基环内向树」上的最长的「双向游走」路径之和
	ring, total := 0, 0
	for i := 0; i < n; i++ {
		if !used[i] {
			j := favorite[i]
			// favorite[favorite[i]] = i 说明环的大小为 2
			if favorite[j] == i {
				total += f[i] + f[j]
				used[i], used[j] = true, true
			} else {
				// 否则环的大小至少为 3，我们需要找出环
				u, cnt := i, 0
				for true {
					cnt++
					u = favorite[u]
					used[u] = true
					if u == i {
						break
					}
				}
				ring = max(ring, cnt)
			}
		}
	}
	return max(ring, total)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// @lc code=end

/*
// @lcpr case=start
// [2,2,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,0]\n
// @lcpr case=end

// @lcpr case=start
// [3,0,1,4,1]\n
// @lcpr case=end

*/

