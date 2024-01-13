class Solution:
    globalSatisfaction = []
    
    def getSum(self, count: int) -> int:
        n = len(globalSatisfaction)
        ans, counter = 0, 1

        for i in range(n - count, n):
            ans += counter * globalSatisfaction[i]
            counter += 1

        return ans

    def maxSatisfaction(self, satisfaction: List[int]) -> int:
        global globalSatisfaction


        n = len(satisfaction)
        satisfaction.sort()
        globalSatisfaction = satisfaction
        l, r, mid = 0, n, 0
        while l < r:
            mid = (l + r + 1) // 2
            if self.getSum(mid) > self.getSum(mid - 1): l = mid
            else: r = mid - 1
        
        return self.getSum(l)