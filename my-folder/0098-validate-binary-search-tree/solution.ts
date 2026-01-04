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

function isValidBST(root: TreeNode | null): boolean {

    let dfs = function(node: TreeNode | null, minBound: number, maxBound: number): boolean {
        if (node === null)
            return true

        if (node.val <= minBound || node.val >= maxBound) {
            return false
        }

        return dfs(node.left, minBound, node.val) && dfs(node.right, node.val, maxBound)

    }

    return dfs(root, -Infinity, Infinity)
    
};
