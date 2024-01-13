class Solution {
public:
    vector<int> globalSatisfaction;
    int getSum(int count) {
        int n = globalSatisfaction.size();
        int counter = 1, ans = 0;
        for (int i = n - count; i <= n - 1; i++) { 
            ans += counter * globalSatisfaction[i];
            counter++;
        }
        return ans;
    }


    int maxSatisfaction(vector<int>& satisfaction) {
        int n = satisfaction.size();
        sort(satisfaction.begin(), satisfaction.end());
        globalSatisfaction = satisfaction;

        int l = 0, r = n, mid;
        while (l < r) {
            mid = (l + r + 1) / 2;
            if (getSum(mid) > getSum(mid - 1)) l = mid;
            else r = mid - 1;
        }

        return getSum(l);
    }
};