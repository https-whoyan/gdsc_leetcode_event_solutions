class Solution:
    def longestAlternatingSubarray(self, nums: List[int], threshold: int) -> int:
        n: int = len(nums)
        for arrSz in range(n, 0, -1):
            for startArr in range(0, n - arrSz + 1):
                isCorrect: bool = True
                if nums[startArr] % 2 != 0:
                    isCorrect = False
                    continue
                for i in range(startArr, arrSz + startArr - 1):
                    if nums[i] % 2 == nums[i + 1] % 2 or nums[i] > threshold:
                        isCorrect = False
                        break
                if nums[arrSz + startArr - 1] > threshold: isCorrect = False
                if isCorrect: return arrSz
        return 0