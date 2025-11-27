class Solution {
public:
    vector<int> relativeSortArray(vector<int>& arr1, vector<int>& arr2) {
          map<int, int> freq;
        for (int x : arr1) freq[x]++;

        vector<int> result;
        for (int x : arr2) {
            while (freq[x] > 0) {
                result.push_back(x);
                freq[x]--;
            }
        }
        for (auto &p : freq) {
            while (p.second > 0) {
                result.push_back(p.first);
                p.second--;
            }
        }
        return result;  
    }
};