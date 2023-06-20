package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxSumBST(root *TreeNode) int {
	// return max sum of a subtree that is also a binary search tree
	result, _, _, _, _ := dfs(root)
	// empty subtree with sum 0 is a valid solution
	return max(0, result)
}

// returns result, sum, min, max and whether subtree rooted at node is a BST
func dfs(node *TreeNode) (int, int, int, int, bool) {
	if node.Left == nil && node.Right == nil {
		return node.Val, node.Val, node.Val, node.Val, true
	}

	sum := node.Val
	mi, ma := node.Val, node.Val
	result := 0
	resultSet := false
	isBST := true
	if node.Left != nil {
		lresult, lsum, lmin, lmax, ok := dfs(node.Left)
		if !ok {
			isBST = false
		} else if lmax >= node.Val {
			isBST = false
		}
		result = lresult
		resultSet = true
		mi = min(mi, lmin)
		ma = max(ma, lmax)
		sum += lsum
	}
	if node.Right != nil {
		rresult, rsum, rmin, rmax, ok := dfs(node.Right)
		if !ok {
			isBST = false
		} else if rmin <= node.Val {
			isBST = false
		}
		if resultSet {
			result = max(result, rresult)
		} else {
			result = rresult
			resultSet = true
		}
		mi = min(mi, rmin)
		ma = max(ma, rmax)
		sum += rsum
	}
	if isBST {
		result = max(result, sum)
	}
	return result, sum, mi, ma, isBST
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
