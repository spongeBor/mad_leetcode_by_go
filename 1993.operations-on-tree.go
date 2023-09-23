package lc

/*
 * @lc app=leetcode.cn id=1993 lang=golang
 * @lcpr version=21917
 *
 * [1993] 树上的操作
 */

// @lc code=start
type LockingTree struct {
	parent       []int
	lockNodeUser []int
	children     [][]int
}

func Constructor(parent []int) LockingTree {
	n := len(parent)
	lockNodeUser := make([]int, n)
	children := make([][]int, n)
	for i := 0; i < n; i++ {
		lockNodeUser[i] = -1
		p := parent[i]
		if p != -1 {
			children[p] = append(children[p], i)
		}
	}
	return LockingTree{parent, lockNodeUser, children}
}

func (this *LockingTree) Lock(num int, user int) bool {
	if this.lockNodeUser[num] == -1 {
		this.lockNodeUser[num] = user
		return true
	}
	return false
}

func (this *LockingTree) Unlock(num int, user int) bool {
	if this.lockNodeUser[num] == user {
		this.lockNodeUser[num] = -1
		return true
	}
	return false
}

func (this *LockingTree) Upgrade(num int, user int) bool {
	res := this.lockNodeUser[num] == -1 && !this.hasLockedAncestor(num) && this.checkAndUnlockDescendant(num)
	if res {
		this.lockNodeUser[num] = user
	}
	return res
}

func (this *LockingTree) hasLockedAncestor(num int) bool {
	num = this.parent[num]
	for num != -1 {
		if this.lockNodeUser[num] != -1 {
			return true
		}
		num = this.parent[num]
	}
	return false
}

func (this *LockingTree) checkAndUnlockDescendant(num int) bool {
	res := false
	if this.lockNodeUser[num] != -1 {
		res = true
	}
	this.lockNodeUser[num] = -1
	for _, child := range this.children[num] {
		if this.checkAndUnlockDescendant(child) {
			res = true
		}
	}
	return res
}

/**
 * Your LockingTree object will be instantiated and called as such:
 * obj := Constructor(parent);
 * param_1 := obj.Lock(num,user);
 * param_2 := obj.Unlock(num,user);
 * param_3 := obj.Upgrade(num,user);
 */
// @lc code=end

/*
// @lcpr case=start
// ["LockingTree", "lock", "unlock", "unlock", "lock", "upgrade", "lock"][[[-1, 0, 0, 1, 1, 2, 2]], [2, 2], [2, 3], [2, 2], [4, 5], [0, 1], [0, 1]]\n
// @lcpr case=end

*/
