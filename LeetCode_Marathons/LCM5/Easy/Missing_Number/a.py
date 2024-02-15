class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        needSumOfNums, sumOfNums = 0, 0
        n = len(nums)

        needSumOfNums = n * (n + 1) // 2
        sumOfNums = sum(nums)

        return needSumOfNums - sumOfNums