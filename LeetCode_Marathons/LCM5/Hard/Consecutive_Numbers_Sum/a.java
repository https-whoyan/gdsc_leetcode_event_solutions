public class Solution {
    private ArrayList<Integer> getAllDivisors(int n) {
        int sqrtN = (int) Math.sqrt(n);
        ArrayList<Integer> allDivisors = new ArrayList<>();
        for (int i = 1; i <= sqrtN; i++) {
            if (n % i == 0) {
                allDivisors.add(i);
                if (n / i != i) {
                    allDivisors.add(n / i);
                }
            }
        }
        return allDivisors;
    }

    public int consecutiveNumbersSum(int n) {
        ArrayList<Integer> allDivisors = getAllDivisors(n);
        int counter = 0;
        for (int divisor : allDivisors) {
            if (divisor % 2 == 1) {
                counter++;
            }
        }
        return counter;
    }
}