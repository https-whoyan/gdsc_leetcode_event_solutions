class Solution:
    def getdDominantEl(self, arr: List[int]) -> int:
        copyArr: List[int] = sorted(arr.copy())
        return copyArr[len(arr) // 2]

    def minimumIndex(self, nums: List[int]) -> int:
        domEl: int = self.getdDominantEl(nums)
        n: int = len(nums)
        pref: List[int] = [0] * n
        postf: List[int] = [0] * n
        for i in range(n):
            if i != 0: pref[i] += pref[i - 1]
            if nums[i] == domEl: pref[i] += 1
        for i in range(n - 1, -1, -1):
            if i != n - 1: postf[i] += postf[i + 1]
            if nums[i] == domEl: postf[i] += 1
        for i in range(n - 1):
            lenFMsv: int = i + 1
            lenSMsv: int = n - lenFMsv
            if pref[i] * 2 > lenFMsv and postf[i + 1] * 2 > lenSMsv:
                return i
        return -1