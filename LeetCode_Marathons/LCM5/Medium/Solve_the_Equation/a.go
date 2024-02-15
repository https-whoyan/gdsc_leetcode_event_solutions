func getXNumsCount(equation string) (int, int) {
    equation += "+"
    intBuilder := 0
    minusOne := 1
    сreadtedBuilder := false
    xCounter, numCounter := 0, 0
    
    for _, token := range equation {
        switch token {
        case '+':
            numCounter += minusOne * intBuilder
            minusOne = 1
            intBuilder = 0
            сreadtedBuilder = false
        case '-':
            numCounter += minusOne * intBuilder
            minusOne = -1
            intBuilder = 0
            сreadtedBuilder = false
        case 'x':
            if intBuilder == 0 && !сreadtedBuilder {
                intBuilder = 1
            }
            xCounter += intBuilder * minusOne
            intBuilder = 0
            сreadtedBuilder = false
        default:
            сreadtedBuilder = true
            intBuilder *= 10
            intBuilder += int(token - '0')
        }
    }
    return xCounter, numCounter
}


func solveEquation(equation string) string {
    leftEquation := strings.Split(equation, "=")[0]
    rightEquation := strings.Split(equation, "=")[1]
    xLCounter, numLCounter := getXNumsCount(leftEquation)
    xRCounter, numRCounter := getXNumsCount(rightEquation)

    xCounter := xLCounter - xRCounter
    numCounter := numRCounter - numLCounter
    if xCounter == 0 {
        if numCounter == 0 {
            return "Infinite solutions"
        } else {
            return "No solution"
        }
    }
    
    xIsEquals := numCounter / xCounter
    return "x=" + strconv.Itoa(xIsEquals)
}