func convertToBase7(num int) string {
    needAddMinusOne := num < 0
    if num == 0 {
        return "0"
    }
    if needAddMinusOne {
        num = -1 * num
    }
    var base7s string
    for num != 0 {
        base7s = string(byte('0' + num % 7)) + base7s
        num /= 7
    }
    if needAddMinusOne {
        return "-" + base7s
    }
    return base7s
}