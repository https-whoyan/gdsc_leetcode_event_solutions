func getdDominantEl(arr []int) int {
    copyArr := make([]int, len(arr))
    copy(copyArr, arr)
    sort.Ints(copyArr)
    return copyArr[len(arr) / 2]
}

func minimumIndex(nums []int) int {
    domEl := getdDominantEl(nums)
    fmt.Println(domEl)
    n := len(nums)
    pref := make([]int, n)
    postf := make([]int, n)
    for i := 0; i <= n - 1; i++ {
        pref[i] = 0
        if i != 0 {
            pref[i] += pref[i - 1]
        }
        if nums[i] == domEl {
            pref[i] += 1
        }
    }
    for i := n - 1; i >= 0; i-- {
        postf[i] = 0
        if i != n - 1 {
            postf[i] += postf[i + 1]
        }
        if nums[i] == domEl {
            postf[i] += 1
        }
    }
    fmt.Println()
    for i := 0; i < n - 1; i++ {
        lenFMsv := i + 1
        lenSMsv := n - lenFMsv
        if pref[i] * 2 > lenFMsv && postf[i + 1] * 2 > lenSMsv {
            return i
        }
    }
    return -1
}