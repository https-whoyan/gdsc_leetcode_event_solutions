class Solution:
    globalGrid = []
    
    def dfs(self, x: int, y: int) -> None:
        globalGrid[x][y] = "0"
        n, m = len(globalGrid), len(globalGrid[0])
        dirr = [[1, 0], [-1, 0], [0, 1], [0, -1]]
        for curDir in dirr:
            nextX, nextY = x + curDir[0], y + curDir[1]
            if not nextX in {-1, n} and not nextY in {-1, m} and globalGrid[nextX][nextY] == "1":
                self.dfs(nextX, nextY)


    def numIslands(self, grid: List[List[str]]) -> int:
        global globalGrid

        globalGrid = grid
        n, m = len(grid), len(grid[0])

        counter = 0

        for i in range(n):
            for j in range(m):
                if globalGrid[i][j] == "1":
                    self.dfs(i, j)
                    counter += 1

        return counter