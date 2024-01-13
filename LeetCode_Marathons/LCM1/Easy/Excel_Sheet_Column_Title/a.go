func convertToTitle(columnNumber int) string {
    var ans string

    for columnNumber != 0 {
        columnNumber--;
        ans = string( 'A' + rune(columnNumber % 26) ) + ans
        columnNumber /= 26
    }

    return ans
}