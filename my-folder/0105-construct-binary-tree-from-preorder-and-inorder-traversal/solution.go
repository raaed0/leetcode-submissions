/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }
    rootVal := preorder[0]
    root := &TreeNode{Val: rootVal}
    index := indexOf(inorder, rootVal)
    if index == -1 {
        return nil
    }

    root.Left = buildTree(preorder[1:index+1], inorder[:index])
    root.Right = buildTree(preorder[index+1:], inorder[index+1:])

    return root
}

func indexOf(slice []int, target int) int {
    for i, v := range slice {
        if v == target {
            return i
        }
    }

    return -1
}
