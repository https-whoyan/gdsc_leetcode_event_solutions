class Solution {
public:
    int maxCount(int m, int n, vector<vector<int>>& ops) {
        int minX = m, minY = n;

        int opsLen = ops.size();
        for (int i = 0; i <= opsLen - 1; i++) {
            minX = min(minX, ops[i][0]);
            minY = min(minY, ops[i][1]);
        }

        return minX * minY;
    }
};