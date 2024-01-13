class Solution:
    def search(self, nums: List[int], target: int) -> int:
        n = len(nums)
        l, r, mid = [0] * 3

        secondArrayStartIndex = 0

        if nums[0] > nums[n - 1]: 
            l, r = 0, n - 1
            while l < r:
                mid = (l + r + 1) // 2
                if nums[mid] > nums[l]: l = mid 
                else:
                    if nums[mid - 1] > nums[mid]: l, r = mid, mid
                    else: r = mid - 1
            secondArrayStartIndex = l
        
        print(secondArrayStartIndex)
        l, r = 0, secondArrayStartIndex - 1
        while l <= r:
            mid = (l + r) // 2
            if nums[mid] == target: return mid
            elif nums[mid] < target: l = mid + 1
            else: r = mid - 1

        l, r = secondArrayStartIndex, n - 1
        while l <= r:
            mid = (l + r) // 2
            if nums[mid] == target: return mid
            elif nums[mid] < target: l = mid + 1
            else: r = mid - 1

        return -1