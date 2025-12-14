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

function maxDepth(root: TreeNode | null): number {
    if (root == null)
        return 0

    let depth = 0
    let queue: Queue<TreeNode> = new Queue()
    queue.enqueue(root)

    while (queue.size() > 0) {
        depth++
        let size = queue.size()
        for (let i=0; i<size; i++) {
            let curr = queue.dequeue()
            if (curr.left != null) {
                queue.enqueue(curr.left)
            }
            if (curr.right != null) {
                queue.enqueue(curr.right)
            }
        }
    }
    
    return depth
};
