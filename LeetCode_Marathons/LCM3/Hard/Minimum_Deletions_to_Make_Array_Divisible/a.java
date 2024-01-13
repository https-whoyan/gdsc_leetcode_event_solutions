class Solution {
    private int gcd(int number1, int number2) {
        while (number2 != 0) {
            int RemainderOfDivision = number1 % number2;
            number1 = number2;
            number2 = RemainderOfDivision;
        }
        return number1;
    }

    public int minOperations(int[] nums, int[] numsDivide) {
        int gcdOfNumsDivine = numsDivide[0];
        for (int number : numsDivide) {
            gcdOfNumsDivine = gcd(gcdOfNumsDivine, number);
        }
        Arrays.sort(nums);
        int counter = 0;
        for (int number : nums) {
            if (gcdOfNumsDivine % number == 0) {
                return counter;
            }
            counter++;
        }
        return -1;
    }
}