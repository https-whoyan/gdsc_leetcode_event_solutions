class Solution:
    def maxCount(self, m: int, n: int, ops: List[List[int]]) -> int:
        minX: int = m
        minY: int = n

        for opInfo in ops:
            minX = min(minX, opInfo[0])
            minY = min(minY, opInfo[1])
        
        return minX * minY