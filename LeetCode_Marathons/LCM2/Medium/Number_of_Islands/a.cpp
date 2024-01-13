class Solution {
public:
    vector<vector<char>> globalGrid;

    void dfs(int x, int y) {
        globalGrid[x][y] = '0';
        int n = globalGrid.size(), m = globalGrid[0].size();
        int dir[4][2] = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};

        for (int dirrIndex = 0; dirrIndex < 4; dirrIndex++) {
            int nextX = x + dir[dirrIndex][0], nextY = y + dir[dirrIndex][1];
            if ((nextX != -1 && nextX != n) && (nextY != -1 && nextY != m) && globalGrid[nextX][nextY] == '1') {
                dfs(nextX, nextY);
            }
        }
    }

    int numIslands(vector<vector<char>>& grid) {
        globalGrid = grid;
        int n = grid.size(), m = grid[0].size();
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
};