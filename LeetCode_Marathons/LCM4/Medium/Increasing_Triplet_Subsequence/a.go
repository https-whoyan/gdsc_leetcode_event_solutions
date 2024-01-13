func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func increasingTriplet(nums []int) bool {
    if len(nums) < 3 {
        return false
    }
    n := len(nums)
    pref := make([]int, n)
    pref[1] = nums[0]
    for i := 2; i <= n - 1; i++ {
        pref[i] = min(pref[i - 1], nums[i - 1])
    }
    postf := make([]int, n)
    postf[n - 2] = nums[n - 1]
    for i := n - 3; i >= 0; i-- {
        postf[i] = max(postf[i + 1], nums[i + 1])
    }
    for i := 1; i <= n - 2; i++ {
        if pref[i] < nums[i] && nums[i] < postf[i] {
            return true
        }
    }
    return false
}