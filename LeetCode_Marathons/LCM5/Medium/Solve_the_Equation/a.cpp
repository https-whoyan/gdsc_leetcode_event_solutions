class Solution {
public:
    vector<int> getXNumsCount(string equation) {
        equation += "+";
        int intBuilder = 0, minusOne = 1;
        int xCounter = 0, numCounter = 0;
        bool сreadtedBuilder = false;

        for (char token : equation) {
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

        return vector<int>{xCounter, numCounter};
    }

    string solveEquation(string equation) {
        string leftStr = equation.substr(0, equation.find('='));
        string rightStr = equation.substr(equation.find('=') + 1);
        vector<int> lValues = getXNumsCount(leftStr);
        vector<int> rValues = getXNumsCount(rightStr);

        int xCounter = lValues[0] - rValues[0];
        int numCounter = rValues[1] - lValues[1];
        if (xCounter == 0) {
            if (numCounter == 0) return "Infinite solutions";
            else return "No solution";
        }

        string xIsEquals = to_string(numCounter / xCounter);
        return "x=" + xIsEquals;
    }
};