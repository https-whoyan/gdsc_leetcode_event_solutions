func getCommon(nums1 []int, nums2 []int) int {
    mpNums2 := make(map[int]bool)
    for _, number2 := range nums2 {
        mpNums2[number2] = true
    }
    for _, number1 := range nums1 {
        _, exists := mpNums2[number1]
        if exists {
            return number1
        }
    }
    return -1
}