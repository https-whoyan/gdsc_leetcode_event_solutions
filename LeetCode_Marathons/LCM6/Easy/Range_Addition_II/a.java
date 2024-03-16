class Solution {
    public int maxCount(int m, int n, int[][] ops) {
        int minX = m, minY = n;

        int opsLen = ops.length;
        for (int i = 0; i <= opsLen - 1; i++) {
            if (ops[i][0] < minX) minX = ops[i][0];
            if (ops[i][1] < minY) minY = ops[i][1];
        }

        return minX * minY;
    }
}