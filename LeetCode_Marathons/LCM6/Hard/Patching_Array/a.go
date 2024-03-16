func minPatches(nums []int, n int) int {
    dpSum := 0
    ans := 0
    nums = append(nums, n + 1)
    for _, num := range nums {
        for dpSum + 1 < num {
            if dpSum >= n {
                return ans
            }
            ans++
            dpSum = dpSum * 2 + 1
        }
        dpSum += num
        if dpSum >= n {
            return ans
        }
    }
    return ans
}