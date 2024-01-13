func countTime(time string) int {
    myHour, myMinutes := time[:2], time[3:]
    myHourAns, myMinutesAns := 1, 1
    fmt.Println(myHour, myMinutes)

    if myHour[1] == byte('?') {
        if myHour[0] == byte('?') {
            myHourAns = 24
        } else if myHour[0] == byte('2') {
            myHourAns = 4
        } else {
            myHourAns = 10
        }
    } else {
        if myHour[0] == byte('?') {
            myHourAns = 2
            if myHour[1] <= byte('3') {
                myHourAns += 1
            }
        }
    }

    if myMinutes[0] == byte('?') {
        myMinutesAns *= 6
    }
    if myMinutes[1] == byte('?') {
        myMinutesAns *= 10
    }

    return myMinutesAns * myHourAns;
}