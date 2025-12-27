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

function isSubtree(root: TreeNode | null, subRoot: TreeNode | null): boolean {
    if (subRoot == null)
        return true
    if (root == null)
        return false

        
    if (sameTree(root, subRoot))
        return true

    return (isSubtree(root.left, subRoot) || isSubtree(root.right, subRoot))
        
};

function sameTree(node: TreeNode | null, subRoot: TreeNode | null): boolean {
    if (node == null && subRoot == null)
        return true

    if (node != null && subRoot != null && node.val == subRoot.val)
        return (sameTree(node.left, subRoot.left) && sameTree(node.right, subRoot.right))

    return false
}
