/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	left := inorderTraversal(root.Left)
	result = append(result, left...)

	result = append(result, root.Val)

	right := inorderTraversal(root.Right)
	result = append(result, right...)

	return result
}
