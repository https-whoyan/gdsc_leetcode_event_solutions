var globalGrid [][]byte;

func dfs(x int, y int) {
    globalGrid[x][y] = byte('0')
    n, m := len(globalGrid), len(globalGrid[0])
    var dir = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

    for dirrIndex := 0; dirrIndex < 4; dirrIndex++ {
        nextX, nextY := x + dir[dirrIndex][0], y + dir[dirrIndex][1]
        if (nextX != -1 && nextX != n) && (nextY != -1 && nextY != m) && globalGrid[nextX][nextY] == byte('1') {
            dfs(nextX, nextY);
        }
    }
}

func numIslands(grid [][]byte) int {
    globalGrid = grid
    n, m := len(grid), len(grid[0])
    counter := 0
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if globalGrid[i][j] == byte('1') {
                dfs(i, j)
                counter++
            }
        }
    }

    return counter
}