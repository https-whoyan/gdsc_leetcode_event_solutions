func twoSum(numbers []int, target int) []int {
    n := len(numbers)
    l, r := 0, n - 1;

    for l < r {
        if numbers[l] + numbers[r] == target {
            return []int{l + 1, r + 1}
        } else if numbers[l] + numbers[r] < target {
            l++;
        } else {
            r--;
        }
    }

    return []int{}
}