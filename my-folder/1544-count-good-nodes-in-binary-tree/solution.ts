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

function goodNodes(root: TreeNode | null): number {
    let goodnodes: number = 0

    let dfs = function(node: TreeNode | null, maxVal: number) {
        if (node === null) {
            return
        }

        if (node.val >= maxVal) {
            goodnodes++
        }

        dfs(node.left, Math.max(node.val, maxVal))
        dfs(node.right, Math.max(node.val, maxVal))

    }
    dfs(root, root.val)
    return goodnodes
};
