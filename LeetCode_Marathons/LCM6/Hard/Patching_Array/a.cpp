class Solution {
public:
    int minPatches(vector<int>& nums, int n) {
        long long dpSum = 0;
        int counter = 0;
        nums.push_back(n);
        
        for (int num: nums) {
            while (dpSum + 1 < num) {
                if (dpSum >= n) return counter;
                counter++;
                dpSum = dpSum + 1 + dpSum;
            }
            dpSum += num;
            if (dpSum >= n) return counter;
        }
        return counter;
    }
};