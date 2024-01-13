class Solution {
    private char[][] globalGrid;

    private void dfs(int x, int y) {
        globalGrid[x][y] = '0';
        int n = globalGrid.length, m = globalGrid[0].length;
        int[][] dir = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};

        for (int dirrIndex = 0; dirrIndex < 4; dirrIndex++) {
            int nextX = x + dir[dirrIndex][0], nextY = y + dir[dirrIndex][1];
            if ((nextX != -1 && nextX != n) && (nextY != -1 && nextY != m)) {
                if (globalGrid[nextX][nextY] == '1') {
                    dfs(nextX, nextY);
                }
            }
        }
    }

    public int numIslands(char[][] grid) {
        globalGrid = grid;

        int n = grid.length, m = grid[0].length;
        int counter = 0;

        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                if (globalGrid[i][j] == '1') {
                    dfs(i, j);
                    counter++;
                }
            }
        }

        return counter;
    }
}