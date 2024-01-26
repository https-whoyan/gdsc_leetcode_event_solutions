class Solution:
    def getCommon(self, nums1: List[int], nums2: List[int]) -> int:
        set1: Set(int) = set(nums1)
        set2: Set(int) = set(nums2)
        commonEls: Set(int) = set.intersection(set1, set2)
        if len(commonEls) == 0: return -1
        else: return sorted(list(commonEls))[0]