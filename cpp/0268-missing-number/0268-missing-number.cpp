class Solution {
public:
    int missingNumber(vector<int>& nums) {
        int n = nums.size();
        int expected_sum = n * (n + 1) / 2;
        
        int actual_sum = 0;
        for (int x : nums) {
            actual_sum += x;
        }
        
        return expected_sum - actual_sum;
    }
};