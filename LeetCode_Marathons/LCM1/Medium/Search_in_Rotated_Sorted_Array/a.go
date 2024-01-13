func search(nums []int, target int) int {
    var (
        n int
        l, r, mid int
    )

    n = len(nums)
    secondArrayStartIndex := 0

    if nums[0] > nums[n - 1] {
        l, r = 0, n - 1
        for l < r {
            mid = (l + r + 1) / 2
            if nums[mid] > nums[l] {
                l = mid
            } else {
                if nums[mid - 1] > nums[mid] {
                    l, r = mid, mid
                } else {
                    r = mid - 1
                }
            }
        }
        secondArrayStartIndex = l
    }

    l, r = 0, secondArrayStartIndex - 1
    for l <= r {
        mid = (l + r + 1) / 2;
        if nums[mid] == target {
            return mid;
        } else if nums[mid] < target {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }

    l, r = secondArrayStartIndex, n - 1
    for l <= r {
        mid = (l + r + 1) / 2;
        if nums[mid] == target {
            return mid;
        } else if nums[mid] < target {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return -1;
}