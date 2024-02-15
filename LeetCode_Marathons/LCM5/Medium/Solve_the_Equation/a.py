class Solution:
    def getXNumsCount(self, equation: str) -> List[int]:
        equation += "+"
        intBuilder: int = 0
        minusOne: int = 1
        сreadtedBuilder: bool = False
        xCounter, numCounter = 0, 0

        for token in equation:
            if token == "+":
                numCounter += minusOne * intBuilder
                minusOne = 1
                intBuilder = 0
                сreadtedBuilder = False
            elif token == "-":
                numCounter += minusOne * intBuilder
                minusOne = -1
                intBuilder = 0
                сreadtedBuilder = False
            elif token == 'x':
                if intBuilder == 0 and not сreadtedBuilder:
                    intBuilder = 1
                xCounter += minusOne * intBuilder
                intBuilder = 0
                сreadtedBuilder = False
            else:
                сreadtedBuilder = True
                intBuilder *= 10
                intBuilder += ord(token) - ord('0') 
        
        return [xCounter, numCounter]


    def solveEquation(self, equation: str) -> str:
        leftEquasion, rightEquasion = equation.split("=")
        xLCounter, numLCounter = self.getXNumsCount(leftEquasion)
        xRCounter, numRCounter = self.getXNumsCount(rightEquasion)

        xCounter: int = xLCounter - xRCounter
        numCounter: int = numRCounter - numLCounter
        if xCounter == 0:
            if numCounter == 0: return "Infinite solutions"
            else: return "No solution"
        
        xIsEquals: int = numCounter // xCounter
        return "x=" + str(xIsEquals)