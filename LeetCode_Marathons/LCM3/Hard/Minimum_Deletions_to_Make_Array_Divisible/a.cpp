class Solution {
public:
    int minOperations(vector<int>& nums, vector<int>& numsDivide) {
        int gcdOfNumsDivine = numsDivide[0];
        for (auto number : numsDivide) {
            gcdOfNumsDivine = gcd(gcdOfNumsDivine, number);
        }
        sort(nums.begin(), nums.end());
        int counter = 0;
        for (auto number : nums) {
            if (gcdOfNumsDivine % number == 0) {
                return counter;
            }
            counter++;
        }
        return -1;
    }
};