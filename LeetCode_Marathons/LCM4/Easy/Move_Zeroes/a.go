func moveZeroes(nums []int)  {
    n := len(nums)
    collision := 0
    for i := 0; i <= n - 1; i++ {
        if nums[i] != 0 {
            nums[i - collision] = nums[i]
            if collision != 0 {
                nums[i] = 0
            }
        } else {
            collision++
        }
    }
}