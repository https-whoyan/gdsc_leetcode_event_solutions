class Solution {
public:
    int getCommon(vector<int>& nums1, vector<int>& nums2) {
        set<int> s1, s2;
        for (int number1 : nums1) {
            s1.insert(number1);
        }
        for (int number2 : nums2) {
            s2.insert(number2);
        }

        set<int> commonEls;
        set_intersection(
            s1.begin(), s1.end(),
            s2.begin(), s2.end(),
            inserter(commonEls, commonEls.begin())
        );
        if (commonEls.size() == 0) return -1;
        else return *(commonEls.begin());
    }
};