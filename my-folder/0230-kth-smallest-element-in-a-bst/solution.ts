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

function kthSmallest(root: TreeNode | null, k: number): number {
    let arr: number[] = []

    let dfs = function(node: TreeNode | null) {
        if (node === null)
            return

        dfs(node.left)
        arr.push(node.val)
        dfs(node.right)
    }

    dfs(root)
    return arr[k-1]
};
