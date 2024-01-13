class Solution:
    def majorityElement(self, arr: List[int]) -> int:
        return sorted(arr)[len(arr) // 2]