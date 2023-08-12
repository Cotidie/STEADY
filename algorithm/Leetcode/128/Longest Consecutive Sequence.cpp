#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        
    }
};

int main() {
    Solution s;
    vector<int> nums;

    // 4
    nums = {100,4,200,1,3,2};
    cout << s.longestConsecutive(nums) << endl;

    // 9
    nums = {0,3,7,2,5,8,4,6,0,1};
    cout << s.longestConsecutive(nums) << endl;
}