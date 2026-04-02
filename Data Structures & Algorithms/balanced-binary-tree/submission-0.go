/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    return dfs(root) != -1
}

func dfs(root *TreeNode) int{
	if root == nil {
		return 0
	}

	left := dfs(root.Left)
	if left == -1{
		return -1
	}


	right := dfs(root.Right)
	if right == -1{
		return -1
	}

	if abs(left - right) > 1{
		return -1
	}

	return 1 + max(left, right)
}

func abs(num int) int{
	if num < 0{
		return -num
	}

	return num
}

func max(a,b int) int{
	if a >= b{
		return a
	}

	return b
}