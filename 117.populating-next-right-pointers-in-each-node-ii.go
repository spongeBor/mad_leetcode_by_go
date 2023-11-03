/*
 * @lc app=leetcode.cn id=117 lang=golang
 * @lcpr version=30104
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		for i, node := range tmp {
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,null,7]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/

