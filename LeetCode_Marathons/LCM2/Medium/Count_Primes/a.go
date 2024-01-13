func countPrimes(n int) int {
    isNotPrimeStatuses := make([]bool, n)
    counter := 0
    for i := 2; i < n; i++ {
        if !isNotPrimeStatuses[i] {
            if math.Sqrt(float64(n)) >= float64(i) {
                for j := i * i; j < n; j += i {
                    isNotPrimeStatuses[j] = true
                }
            }
            counter++;
        }
    }
    return counter;
}