func getNWithoutDivizor(divizor, n int) int {
    for n % divizor == 0 {
        n /= divizor
    }
    return n
}

func isUgly(n int) bool {
    if n <= 0 {
        return false
    }
    n = getNWithoutDivizor(2, n)
    n = getNWithoutDivizor(3, n)
    n = getNWithoutDivizor(5, n)
    return n == 1 
}