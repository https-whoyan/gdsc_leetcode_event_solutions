class Solution:
    def getNWithoutDivizor(self, divizor: int, n: int) -> int:
        while n % divizor == 0: n /= divizor
        return n

    def isUgly(self, n: int) -> bool:
        if n <= 0: return False
        n = self.getNWithoutDivizor(2, n)
        n = self.getNWithoutDivizor(3, n)
        n = self.getNWithoutDivizor(5, n)
        return n == 1