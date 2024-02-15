func missingNumber(nums []int) int {
    var needSumOfNums, sumOfNums int
    n := len(nums)

    needSumOfNums = n * (n + 1) / 2;
    for _, num := range(nums) {
        sumOfNums += num
    }

    return needSumOfNums - sumOfNums;
}