class Solution {
public:
    int countHillValley(vector<int>& nums) {
                vector<int> a;
        // Remove consecutive duplicates
        for (int x : nums) {
            if (a.empty() || a.back() != x)
                a.push_back(x);
        }

        int count = 0;
        for (int i = 1; i + 1 < a.size(); i++) {
            if ((a[i] > a[i - 1] && a[i] > a[i + 1]) ||
                (a[i] < a[i - 1] && a[i] < a[i + 1])) {
                count++;
            }
        }

        return count;
    }
};