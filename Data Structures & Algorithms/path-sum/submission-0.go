/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
	var dfs func(root *TreeNode, currSum int) bool
	dfs = func(root *TreeNode, currSum int) bool{
		if root == nil{
			return false
		}

		currSum += root.Val

		if root.Left == nil && root.Right == nil{
			return currSum == targetSum
		}

		return dfs(root.Left, currSum) || dfs(root.Right, currSum)

	}

	return dfs(root, 0)

}
