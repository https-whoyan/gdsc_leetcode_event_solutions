class Solution:
    def getCommon(self, nums1: List[int], nums2: List[int]) -> int:
        n2: int = len(nums2)
        for number1 in nums1:
            l, r, mid = 0, n2 - 1, 0
            while l < r:
                mid = (l + r) // 2
                if nums2[mid] == number1: return number1
                elif nums2[mid] < number1: l = mid + 1
                else: r = mid - 1
            if nums2[l] == number1: return number1
        return -1