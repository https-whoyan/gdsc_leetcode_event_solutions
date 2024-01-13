func gcd(number1 int, number2 int) int {
    for number2 != 0 {
        number1, number2 = number2, number1 % number2
    }
    return number1
}

func minOperations(nums []int, numsDivide []int) int {
    gcdOfNumsDivine := numsDivide[0]
    for i := 1; i < len(numsDivide); i++ {
        gcdOfNumsDivine = gcd(gcdOfNumsDivine, numsDivide[i])
    }
    sort.Ints(nums)
    counter := 0
    for _, number := range nums {
        if gcdOfNumsDivine % number == 0 {
            return counter
        }
        counter++
    }
    return -1
}