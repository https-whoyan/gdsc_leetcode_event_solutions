class Solution:
    def maxCount(self, m: int, n: int, ops: List[List[int]]) -> int:
        return min([x[0] for x in ops] + [m]) * min([x[1] for x in ops] + [n])