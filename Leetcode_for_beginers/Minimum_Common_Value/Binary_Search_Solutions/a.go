func getCommon(nums1 []int, nums2 []int) int {
    n2 := len(nums2)
    for _, number1 := range nums1 {
        l, r, mid := 0, n2 - 1, 0
        for l < r {
            mid = (l + r) / 2
            if nums2[mid] == number1 {
                return number1
            } else if nums2[mid] < number1 {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
        if nums2[l] == number1 {
            return number1
        }
    }
    return -1
}