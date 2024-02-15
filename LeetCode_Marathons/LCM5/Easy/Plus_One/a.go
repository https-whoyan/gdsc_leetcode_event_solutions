func plusOne(digits []int) []int {
    n := len(digits)
    digits[n - 1]++
    for i := n - 1; i >= 0; i-- {
        if digits[i] >= 10 {
            digits[i] = 0
            if i != 0 {
                digits[i - 1]++
            } else {
                return append([]int{1}, digits...)
            }
        } else {
            return digits
        }
    }
    return []int{}
}