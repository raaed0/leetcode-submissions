type Point struct {
    X, Y int
}
type State struct {
    R, C, Path int
}

func shortestPathBinaryMatrix(grid [][]int) int {
    n := len(grid)
    if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
        return -1
    }
    queue := []State{}
    queue = append(queue, State{0,0,1})
    directions := [][]int{{1,0}, {0,1}, {-1,0}, {0,-1}, {1,1}, {-1,-1}, {-1,1}, {1,-1}}
    visited := make(map[Point]bool)
    visited[Point{0,0}] = true

    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]

        r, c, path := curr.R, curr.C, curr.Path

        if min(r,c) < 0 || max(r,c) == n || grid[r][c] == 1 {
            continue
        }

        if r == n-1 && c == n-1 {
            return path
        }

        for _, dir := range directions {
            dr, dc := r+dir[0], c+dir[1]

            if !visited[Point{dr, dc}] && dr >= 0 && dc >= 0 && dr < n && dc < n && grid[dr][dc] == 0 {
                queue = append(queue, State{dr, dc, path+1})
                visited[Point{dr, dc}] = true
            } 
        }
    }

    return -1
}

