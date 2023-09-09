#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

vector<long> numsA = vector<long>(1001, 0);
vector<long> numsB = vector<long>(1001, 0);

int findLeftIdx(long target, vector<long> nums) {
    int low = 0;
    int high = nums.size()-1;

    while (low <= high) {
        int mid = (low + high) / 2;

        if (nums[mid] < target) {
            low = mid+1;
        } else {
            high = mid-1;
        }
    }

    return low;
}

int findRightIdx(long target, vector<long> nums) {
    int low = 0;
    int high = nums.size()-1;

    while (low <= high) {
        int mid = (low + high) / 2;

        if (nums[mid] <= target) {
            low = mid+1;
        } else {
            high = mid-1;
        }
    }

    return low;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(NULL);

    long target;
    cin >> target;

    int n;
    cin >> n;

    for (int i=0; i<n; i++) {
        int val;
        cin >> val;

        numsA[i] = val;
    }

    int m;
    cin >> m;

    for (int i=0; i<m; i++) {
        int val;
        cin >> val;

        numsB[i] = val;
    }

    vector<long> sumsA;
    for (int i=0; i<n; i++) {
        long sum=0;
        for (int j=i; j<n; j++) {
            sum += numsA[j];
            sumsA.push_back(sum);
        }
    }

    vector<long> sumsB;
    for (int i=0; i<m; i++) {
        long sum = 0;
        for (int j=i; j<m; j++) {
            sum += numsB[j];
            sumsB.push_back(sum);
        }
    }
    
    sort(sumsB.begin(), sumsB.end());

    long cnt = 0;
    for (auto sumA : sumsA) {
        int diff = target - sumA;
        
        auto leftIdx = lower_bound(sumsB.begin(), sumsB.end(), diff);
        auto rightIdx = upper_bound(sumsB.begin(), sumsB.end(), diff);

        cnt += (rightIdx - leftIdx);
    }

    cout << cnt << '\n';
}