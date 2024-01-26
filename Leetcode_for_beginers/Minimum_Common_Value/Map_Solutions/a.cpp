class Solution {
public:
    int getCommon(vector<int>& nums1, vector<int>& nums2) {
        map<int, bool> mpNums2;
        for (int number2 : nums2) {
            mpNums2[number2] = true;
        }
        for (int number1 : nums1) {
            if (mpNums2[number1]) {
                return number1;
            }
        }
        return -1;
    }
};