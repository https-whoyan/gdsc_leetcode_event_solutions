class Solution {
public:
    int smallestRangeI(vector<int>& nums, int k) {
        int minNum = INT_MAX, maxNum = INT_MIN;

        for (int num : nums) {
            if (num < minNum) minNum = num;
            if (num > maxNum) maxNum = num;
        }

        if (maxNum - minNum <= 2 * k) return 0;
        return maxNum - minNum - 2 * k;
    }
};