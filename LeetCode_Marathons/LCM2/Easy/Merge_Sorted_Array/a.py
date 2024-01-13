class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        ans_list = []
        id1, id2 = 0, 0
        while id1 + id2 != n + m:
            if id1 == m:
                ans_list.append(nums2[id2])
                id2 += 1
            elif id2 == n or nums1[id1] <= nums2[id2]:
                ans_list.append(nums1[id1])
                id1 += 1
            else:
                ans_list.append(nums2[id2])
                id2 += 1
        
        for i in range(n + m):
            nums1[i] = ans_list[i]