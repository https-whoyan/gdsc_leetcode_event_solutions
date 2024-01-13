class Solution:
    def increasingTriplet(self, nums: List[int]) -> bool:
        value1: int = 2147483647
        value2: int = 2147483647
        for num in nums:
            if num <= value1: value1 = num
            elif num <= value2: value2 = num
            else: return True
        return False