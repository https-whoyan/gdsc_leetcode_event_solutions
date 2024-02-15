class Solution {
public:
    vector<int> plusOne(vector<int>& digits) {
        int n = digits.size();

        digits[n - 1]++;
        for (int i = n - 1; i >= 0; i--) {
            if (digits[i] == 10) {
                digits[i] = 0;
                if (i != 0) digits[i - 1]++;
                else {
                    vector<int> ans {1};
                    for (int j = 0; j <= n - 1; j++) {
                        ans.push_back(digits[j]);
                    }
                    return ans;
                }
            } else return digits;
        }
        return digits;
    }
};