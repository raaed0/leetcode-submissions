/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    res := []int{}

    records := [][]int{}
    
    var dfs func(root *TreeNode, depth int)
    dfs = func(root *TreeNode, depth int) {
        if root == nil {
            return
        }
        if len(records) == depth {
            records = append(records, []int{})
        }
        records[depth] = append(records[depth], root.Val)
        dfs(root.Left, depth+1)
        dfs(root.Right, depth+1)
    }


    dfs(root, 0)

    for _, level := range records {
        res = append(res, level[len(level)-1])
    }

    return res
}
