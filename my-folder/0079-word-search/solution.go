func exist(board [][]byte, word string) bool {
    rows := len(board)
    column := len(board[0])

    visited := make([][]bool, rows)

    for i := range visited {
        visited[i] = make([]bool, column)
    }

    var dfs func(i, j, idx int) bool
    dfs = func(i, j, idx int) bool {
        if idx == len(word) {
            return true
        }

        if i < 0 || j < 0 || i >= rows || j >= column || visited[i][j] || board[i][j] != word[idx]  {
            return false
        }

        visited[i][j] = true

        if dfs(i+1, j, idx+1) || dfs(i-1, j, idx+1) || dfs(i, j+1, idx+1) || dfs(i, j-1, idx+1) {
            return true
        }
        
        visited[i][j] = false
        return false
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < column; j++ {
            if dfs(i, j, 0) {
                return true
            }
        }
    }

    return false
}
