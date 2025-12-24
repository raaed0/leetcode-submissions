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

function isBalanced(root: TreeNode | null): boolean {

    let checkHeight = function(root: TreeNode | null): number {
        if (root == null) {
            return 0
        }

        let left = checkHeight(root.left)
        let right = checkHeight(root.right)

        if (left === -1 || right === -1) {
            return -1
        }
        if (Math.abs(left - right) > 1) {
            return -1
        }

        return 1 + Math.max(left, right)
    }

    return checkHeight(root) !== -1
};
