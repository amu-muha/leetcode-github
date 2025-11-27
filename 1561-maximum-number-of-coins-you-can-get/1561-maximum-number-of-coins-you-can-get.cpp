class Solution {
public:
    int maxCoins(vector<int>& piles) {
            sort(piles.rbegin(), piles.rend());
        int n = piles.size() / 3;
        int res = 0;
        for (int i = 1; i < 2 * n; i += 2) {
            res += piles[i];
        }
        
        return res;
    }
};