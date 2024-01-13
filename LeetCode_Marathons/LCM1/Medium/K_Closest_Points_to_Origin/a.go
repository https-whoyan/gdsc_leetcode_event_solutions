func kClosest(points [][]int, k int) [][]int {
    n := len(points)
    var listDistance [][]int
    for i := 0; i < n; i++ {
        distance := points[i][0] * points[i][0] + points[i][1] * points[i][1]
        listDistance = append(listDistance, []int{distance, i})
    }

    sort.Slice(listDistance, func(i, j int) bool {
        return listDistance[i][0] < listDistance[j][0]
    })

    var ans[][]int
    for i := 0; i < k; i++ {
        ans = append(ans, points[listDistance[i][1]])
    }

    return ans
}