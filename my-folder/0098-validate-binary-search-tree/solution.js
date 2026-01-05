/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {boolean}
 */
var isValidBST = function(root) {
    
    let dfs = function(node, minBound, maxBound) {
        if (node === null)
            return true

        if (node.val <= minBound || node.val >= maxBound) {
            return false
        }

        return dfs(node.left, minBound, node.val) && dfs(node.right, node.val, maxBound)
    }

    return dfs(root, -Infinity, Infinity)
};
