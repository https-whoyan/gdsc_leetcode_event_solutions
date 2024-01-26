class Solution:
    def getCommon(self, nums1: List[int], nums2: List[int]) -> int:
        mpNums2: Dict[int, bool] = defaultdict(bool)
        for number2 in nums2:
            mpNums2[number2] = True
        for number1 in nums1:
            if mpNums2[number1]: return number1
        return -1