func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    memo := make([][]int, m)

    for i := range memo {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(int, int) int
    dfs = func(i, j int) int {
        if i == m {
            return n-j
        }
        if j == n {
            return m-i
        }

        if memo[i][j] != -1 {
            return memo[i][j]
        }

        if word1[i] == word2[j] {
            memo[i][j] = dfs(i + 1, j + 1)
        } else {
            inserted := dfs(i, j+1)
            deleted := dfs(i + 1, j)
            replaced := dfs(i+1, j+1)
            memo[i][j] = 1 + min(inserted, min(deleted, replaced))
        }
        
        return memo[i][j]
    }

    return dfs(0, 0)
}
