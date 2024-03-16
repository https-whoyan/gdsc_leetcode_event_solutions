func maxCount(m int, n int, ops [][]int) int {
    minX, minY := m, n;

    for _, opInfo := range ops {
        minX = min(minX, opInfo[0])
        minY = min(minY, opInfo[1])
    }

    return minX * minY
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}