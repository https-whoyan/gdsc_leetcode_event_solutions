class Solution:
    def countHomogenous(self, s: str) -> int:
        mod: int = 1000000007
        ans, keyCounter = 0, 0
        s += "*"
        key = "*"

        for myCh in s:
            if myCh == key: keyCounter += 1
            else:
                ansForKey: int
                ansForKey = (keyCounter + 1) * keyCounter // 2
                ans = (ans + ansForKey) % mod
                keyCounter = 1
                key = myCh
        
        return ans