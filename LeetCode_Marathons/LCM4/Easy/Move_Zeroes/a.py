class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        n: int = len(nums)
        collision: int = 0
        for i in range(n):
            if nums[i] != 0:
                nums[i - collision] = nums[i]
                if collision != 0: nums[i] = 0
            else: collision += 1