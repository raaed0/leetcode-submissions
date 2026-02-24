/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    nodeMap := make(map[*Node]*Node)

    var dfs func(node *Node) *Node
    dfs = func(node *Node) *Node {
        if node == nil {
            return nil
        }
        if _, exists := nodeMap[node]; exists {
            return nodeMap[node]
        }

        newNode := &Node{Val: node.Val}
        nodeMap[node] = newNode

        for _, nbr := range node.Neighbors {
            newNode.Neighbors = append(newNode.Neighbors, dfs(nbr))
        }
        return newNode
    }

    return dfs(node)
}
