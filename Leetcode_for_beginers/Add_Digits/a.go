func addDigits(num int) int {
    for num >= 10 {
        sumOfDigits := 0
        copyNum := num
        for copyNum != 0 {
            sumOfDigits += (copyNum % 10)
            copyNum /= 10
        }
        num = sumOfDigits
    }
    return num
}