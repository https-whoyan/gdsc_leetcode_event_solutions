class Solution {
    private int[] getXNumsCount(String equation) {
        equation += "+";
        int intBuilder = 0, minusOne = 1;
        int xCounter = 0, numCounter = 0;
        Boolean сreadtedBuilder = false;
        
        for (Character token : equation.toCharArray()) {
            if (token == '+') {
                numCounter += minusOne * intBuilder;
                minusOne = 1;
                intBuilder = 0;
                сreadtedBuilder = false;
            } else if (token == '-') {
                numCounter += minusOne * intBuilder;
                minusOne = -1;
                intBuilder = 0;
                сreadtedBuilder = false;
            } else if (token == 'x') {
                if (intBuilder == 0 && !сreadtedBuilder) {
                    intBuilder = 1;
                }
                xCounter += minusOne * intBuilder;
                intBuilder = 0;
                сreadtedBuilder = false;
            } else {
                сreadtedBuilder = true;
                intBuilder *= 10;
                intBuilder += token - '0';
            }
        }

        return new int[]{xCounter, numCounter};
    }
    
    public String solveEquation(String equation) {
        String[] splittedEquation = equation.split("=");
        int[] lValues = getXNumsCount(splittedEquation[0]);
        int[] rValues = getXNumsCount(splittedEquation[1]);
        
        int xCounter = lValues[0] - rValues[0];
        int numCounter = rValues[1] - lValues[1];
        if (xCounter == 0) {
            if (numCounter == 0) return "Infinite solutions";
            else return "No solution";
        }

        String xIsEquals = Integer.toString(numCounter / xCounter);
        return "x=" + xIsEquals;
    }
}