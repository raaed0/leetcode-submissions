/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    var dfs func(root *TreeNode, maxValue int) int
    dfs = func(root *TreeNode, maxValue int) int {
        res := 0
        if root == nil {
            return res
        }

        if root.Val >= maxValue {
            res = 1
        } else {
            res = 0
        }
        maxValue = max(maxValue, root.Val)

        return res + dfs(root.Left, maxValue) + dfs(root.Right, maxValue)
    }

    return dfs(root, root.Val)
}
