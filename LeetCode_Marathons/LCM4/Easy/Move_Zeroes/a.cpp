class Solution {
public:
    void moveZeroes(vector<int>& nums) {
        int n = nums.size();
        int collision = 0;
        for (int i = 0; i <= n - 1; i++) {
            if (nums[i] != 0) {
                nums[i - collision] = nums[i];
                if (collision != 0) {
                    nums[i] = 0;
                }
            } else {
                collision++;
            }
        }
    }
};