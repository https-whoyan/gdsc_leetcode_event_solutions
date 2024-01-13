class Solution:
    def constructArray(self, n: int, k: int) -> List[int]:
        ans: List[int] = []
        l, r = 1, k + 1
        while l <= r:
            ans = ans + [l]
            l += 1
            if l > r: break
            ans = ans + [r]
            r -= 1
        for i in range(k + 2, n + 1): ans = ans + [i]
        return ans