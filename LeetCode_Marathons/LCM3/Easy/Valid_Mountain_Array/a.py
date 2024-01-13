class Solution:
    def validMountainArray(self, arr: List[int]) -> bool:
        n: int = len(arr)
        if n < 3: return False
        isIncreasing: bool = True
        for i in range(n - 1):
            if isIncreasing:
                if arr[i + 1] < arr[i]:
                    if i != 0: isIncreasing = False
                    else: return False
                elif arr[i + 1] == arr[i]: return False
            else:
                if arr[i + 1] >= arr[i]: return False
        return not isIncreasing