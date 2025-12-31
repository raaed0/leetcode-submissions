/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

function rightSideView(root: TreeNode | null): number[] {
    let res: number[] = []
    
    let dfs = function(root: TreeNode | null, depth: number) {
        if (root == null)
            return

        if (res.length === depth) {
            res.push(root.val)
        }
        
        dfs(root.right, depth + 1)
        dfs(root.left, depth + 1)
    }

    dfs(root, 0)
    return res
};
