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

function diameterOfBinaryTree(root: TreeNode | null): number {
    let diameter = 0

    let dfs = function(root: TreeNode | null) {
        if (root == null) {
            return 0
        }

        let left = dfs(root.left)
        let right = dfs(root.right)

        diameter = Math.max(diameter, left+right)

        return 1 + Math.max(left, right)
    }

    dfs(root)

    return diameter
};


