class Solution:
    def minPatches(self, nums: List[int], n: int) -> int:
        counter: int = 0
        dpSum: int = 0
        nums.append(n + 1)
        for num in nums:
            while dpSum + 1 < num:
                if dpSum >= n: return counter
                counter += 1
                dpSum = dpSum + 1 + dpSum
            dpSum += num
            if dpSum >= n: return counter
        return counter