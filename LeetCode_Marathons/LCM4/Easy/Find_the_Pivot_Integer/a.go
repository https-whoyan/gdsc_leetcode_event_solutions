func pivotInteger(n int) int {
    sumN := n * (n + 1) / 2
    dpSum := 0
    for i := 1; i <= n; i++ {
        dpSum += i
        if dpSum == sumN - dpSum + i {
            return i
        }
    }
    return -1
}