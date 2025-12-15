/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    diameter := 0

    dfs(root, &diameter)
    return diameter
}

func dfs(node *TreeNode, diameter *int) int {
    if node == nil {
        return 0
    }
    left := dfs(node.Left, diameter)
    right := dfs(node.Right, diameter)
    *diameter = max(*diameter, left + right)

    return 1 + max(left, right)
}
