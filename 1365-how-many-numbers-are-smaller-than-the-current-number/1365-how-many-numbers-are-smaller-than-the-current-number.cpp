class Solution {
public:
    vector<int> smallerNumbersThanCurrent(vector<int>& nums) {
          vector<int> count(101, 0);

        for (int x : nums) {
            count[x]++;
        }
        for (int i = 1; i < 101; i++) {
            count[i] += count[i - 1];
        }
        vector<int> result;
        result.reserve(nums.size());
        for (int x : nums) {
            if (x == 0) {
                result.push_back(0);
            } else {
                result.push_back(count[x - 1]);
            }
        }

        return result;
    }
};