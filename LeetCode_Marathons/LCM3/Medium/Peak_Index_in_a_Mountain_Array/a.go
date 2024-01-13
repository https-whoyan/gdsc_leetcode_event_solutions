func peakIndexInMountainArray(arr []int) int {
    l, r, mid := 1, len(arr) - 2, 0
    for l < r {
        mid = (l + r) / 2
        if arr[mid - 1] < arr[mid] {
            if arr[mid] < arr[mid + 1] {
                l = mid + 1
            } else {
                return mid
            }
        } else {
            r = mid - 1
        }
    }
    return l
}