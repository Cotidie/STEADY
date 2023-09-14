#include <iostream>
#include <algorithm>
using namespace std;

int main() {
    int num;
    cin >> num;

    vector<int> nums(num);
    for (int i=0; i<num; i++) {
        cin >> nums[i];
    }

    vector<int> LIS;
    vector<int> answer;

    LIS.push_back(nums[0]);
    
    for (int i=1; i<num; i++) {
        if (nums[i] > LIS.back()) {
            LIS.push_back(nums[i]);
            answer = LIS;
        } else {
            int low = lower_bound(LIS.begin(), LIS.end(), nums[i]) - LIS.begin();
            if (low == LIS.size()-1) LIS[low] = nums[i];
        }
    }

    cout << LIS.size() << '\n';
    for (auto ans :answer) cout << ans << " ";
}