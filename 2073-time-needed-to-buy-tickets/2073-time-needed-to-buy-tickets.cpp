class Solution {
public:
    int timeRequiredToBuy(vector<int>& tickets, int k) {
                int time = 0;
        int n = tickets.size();
        
        for (int i = 0; i < n; i++) {
            time += min(tickets[i], tickets[k]);
            if (i > k && tickets[i] >= tickets[k]) {
                time--;
            }
        }
        
        return time;
    }
};