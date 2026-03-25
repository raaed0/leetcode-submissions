impl Solution {
    pub fn min_path_sum(grid: Vec<Vec<i32>>) -> i32 {
        let row: usize = grid.len();
        let col: usize = grid[0].len();

        let mut dp: Vec<i32> = vec![i32::MAX; col + 1];
        dp[col-1] = 0;

        for r in (0..row).rev() {
            for c in (0..col).rev() {
                dp[c] = grid[r][c] + dp[c].min(dp[c+1]);
            }
        }

        return dp[0];
    }
}
