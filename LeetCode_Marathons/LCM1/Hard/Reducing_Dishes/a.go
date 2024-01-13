var globalSatisfaction []int

func getSum(count int) int {
    n := len(globalSatisfaction)
    ans, counter := 0, 1

    for i := n - count; i <= n - 1; i++ {
        ans += counter * globalSatisfaction[i];
        counter++;
    }

    return ans;
}

func maxSatisfaction(satisfaction []int) int {
    n := len(satisfaction)
    sort.Ints(satisfaction)
    globalSatisfaction = satisfaction

    l, r, mid := 0, n, 0
    for l < r {
        mid = (l + r + 1) / 2
        if getSum(mid) > getSum(mid - 1) {
            l = mid
        } else {
            r = mid - 1
        }
    }

    return getSum(l)
}