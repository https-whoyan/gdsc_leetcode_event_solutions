func longestAlternatingSubarray(nums []int, threshold int) int {
    n := len(nums)
    fmt.Println(nums)
    for arrSz := n; arrSz >= 1; arrSz-- {
        for startArr := 0; startArr <= n - arrSz; startArr++ {
            isCorrect := true
            if nums[startArr] % 2 != 0 {
                isCorrect = false
                continue
            }
            for i := startArr; i <= arrSz + startArr - 2; i++ {
                if nums[i] % 2 == nums[i + 1] % 2 || nums[i] > threshold {
                    isCorrect = false
                    break
                }
            }
            if nums[arrSz + startArr - 1] > threshold {
                isCorrect = false
            }
            if isCorrect {
                return arrSz
            }
        }
    }
    return 0
}