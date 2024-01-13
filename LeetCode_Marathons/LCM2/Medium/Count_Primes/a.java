class Solution {
    public int countPrimes(int n) {
        boolean[] isNotPrimeStatuses = new boolean[n + 1];
        int counter = 0;
        for (int i = 2; i < n; i++) {
            if (!isNotPrimeStatuses[i]) {
                if (Math.sqrt(n) >= i) {
                    for (int j = i * i; j < n; j += i) isNotPrimeStatuses[j] = true;
                }
                counter++;
            }
        }
        return counter;
    }
}