function exist(board: string[][], word: string): boolean {
    let visited: boolean[][] = Array(board.length).fill(false).map(() => Array(board[0].length).fill(false))

    let dfs = function(i: number, j: number, idx: number) {
        if (word.length === idx) {
            return true
        }

        if (i < 0 || j < 0 || i >= board.length || j >= board[0].length || visited[i][j] || board[i][j] != word[idx]) {
            return false
        }

        visited[i][j] = true

        if (dfs(i+1, j, idx+1) || dfs(i-1, j, idx+1) || dfs(i, j+1, idx+1) || dfs(i, j-1, idx+1)) {
            return true
        }

        visited[i][j] = false
        return false
    }

    for (let i = 0; i < board.length; i++) {
        for (let j = 0; j < board[i].length; j++) {
            if (dfs(i, j, 0)) {
                return true
            }
        }
    }

    return false
};
