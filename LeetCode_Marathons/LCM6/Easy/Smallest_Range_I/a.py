class Solution:
    def smallestRangeI(self, nums: List[int], k: int) -> int:
        minNum: int = min(nums)
        maxNum: int = max(nums)

        if maxNum - minNum <= 2 * k: return 0
        return (maxNum - minNum) - 2 * k