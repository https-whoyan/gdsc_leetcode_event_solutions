func constructArray(n int, k int) []int {
    l, r := 1, k + 1
    ans := []int{}
    for l <= r {
        ans = append(ans, l)
        l++
        if l > r {
            break
        }
        ans = append(ans, r)
        r--
    }
    for i := k + 2; i <= n; i++ {
        ans = append(ans, i)
    }
    return ans
}