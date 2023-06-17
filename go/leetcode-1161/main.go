package main

func maxLevelSum(root *TreeNode) int {
	levelSum := make(map[int]int)
	dfs(root, 0, levelSum)
	maxSum := -(1 << 60)
	level := -1
	for i, v := range levelSum {
		if v > maxSum {
			maxSum = v
			level = i
		} else if v == maxSum && i < level {
			level = i
		}
	}
	return level + 1
}

func dfs(node *TreeNode, level int, levelSum map[int]int) {
	if node == nil {
		return
	}
	levelSum[level] += node.Val
	dfs(node.Left, level+1, levelSum)
	dfs(node.Right, level+1, levelSum)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
