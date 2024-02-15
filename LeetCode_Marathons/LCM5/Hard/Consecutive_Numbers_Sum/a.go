func getAllDivizors(n int) []int {
    allDivizors := []int{}
    sqrtN := int(math.Sqrt(float64(n)))
    for i := 1; i <= sqrtN; i++ {
        if n % i == 0 {
            allDivizors = append(allDivizors, i)
            if n / i != i {
                allDivizors = append(allDivizors, n / i)
            }
        }
    }
    return allDivizors
}

func consecutiveNumbersSum(n int) int {
    allDivizors := getAllDivizors(n)
    counter := 0
    for _, divizor := range allDivizors {
        if divizor % 2 == 1 {
            counter++
        }
    }
    return counter
}