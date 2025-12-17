/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    
    var dfs func(node *TreeNode) (bool, int)

    dfs = func(node *TreeNode) (bool, int) {
        if node == nil {
            return true, 0
        }
        ln, lv := dfs(node.Left)
        rn, rv := dfs(node.Right)

        if !ln || !rn {
            return false, 0
        }

        diff := abs(lv - rv)
        if diff <= 1 {
            return true, 1 + max(lv, rv)
        } else {
            return false, 1 + max(lv, rv)
        }
    }

    b, _ := dfs(root)

    return b
}

func abs(v int) int {
    if v < 0 {
        return -v
    }
    return v
}
