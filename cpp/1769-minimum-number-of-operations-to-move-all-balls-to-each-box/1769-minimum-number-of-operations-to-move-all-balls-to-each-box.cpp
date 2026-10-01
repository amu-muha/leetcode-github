class Solution {
public:
    vector<int> minOperations(string boxes) {
      int n = boxes.size();
        vector<int> ans(n, 0);

        int balls = 0;    // balls seen so far (left side)
        int moves = 0;    // operations accumulated from left

        // Left to right
        for (int i = 0; i < n; i++) {
            ans[i] += moves;
            if (boxes[i] == '1') balls++;
            moves += balls;
        }

        balls = 0;
        moves = 0;

        // Right to left
        for (int i = n - 1; i >= 0; i--) {
            ans[i] += moves;
            if (boxes[i] == '1') balls++;
            moves += balls;
        }

        return ans   
    }
};