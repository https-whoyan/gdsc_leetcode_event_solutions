func merge(nums1 []int, m int, nums2 []int, n int)  {
    ansMsv := make([]int, 0, m + n)
    id1, id2 := 0, 0

    for id1 + id2 != n + m {
        if (id1 == m) {
            ansMsv = append(ansMsv, nums2[id2])
            id2++
        } else if (id2 == n || nums1[id1] < nums2[id2]) {
            ansMsv = append(ansMsv, nums1[id1])
            id1++
        } else {
            ansMsv = append(ansMsv, nums2[id2])
            id2++
        }
    }

    for i := 0; i < n + m; i++ {
        nums1[i] = ansMsv[i]
    }
}