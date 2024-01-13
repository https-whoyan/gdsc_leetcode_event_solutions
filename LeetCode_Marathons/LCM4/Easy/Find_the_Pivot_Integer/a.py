class Solution:
    def pivotInteger(self, n: int) -> int:
        sumN: int = n * (n + 1) // 2
        dpSum: int = 0
        for i in range(1, n + 1):
            dpSum += i
            if dpSum == sumN - dpSum + i: return i
        return -1