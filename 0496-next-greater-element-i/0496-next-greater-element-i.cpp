class Solution {
public:
    vector<int> nextGreaterElement(vector<int>& nums1, vector<int>& nums2) {
      unordered_map<int, int> nextGreater;
        stack<int> st;
        for (int num : nums2) {
            while (!st.empty() && num > st.top()) {
                nextGreater[st.top()] = num;
                st.pop();
            }
            st.push(num);
        }
        while (!st.empty()) {
            nextGreater[st.top()] = -1;
            st.pop();
        }
        vector<int> ans;
        for (int x : nums1) {
            ans.push_back(nextGreater[x]);
        }

        return ans;   
    }
};