class Solution {
public:
    int missingNumber(vector<int>& nums) {
        int needSumOfNums = 0, sumOfNums = 0;
        int n = nums.size();

        needSumOfNums = n * (n + 1) / 2;
        for (int i = 0; i < n; i++) sumOfNums += nums[i];

        return needSumOfNums - sumOfNums;
    }
};