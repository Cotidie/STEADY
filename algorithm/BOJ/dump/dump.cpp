#include <iostream>
#include <vector>

using namespace std;

const int INTMAX = 2001;
vector<vector<int>> constraints = {
    {0, 1}, {0, 2},
    {1, 0}, {1, 2},
    {2, 0}, {2, 1},
};

int main() {
    int size;
    cin >> size;

    vector<vector<int>> costs = vector<vector<int>>(
        size+1, vector<int>(
            3, 0
        )
    );

    for (int i=1; i<=size; i++) {
        for (int j=0; j<3; j++) {
            int cost;
            cin >> cost;

            costs[i][j] = cost;
        }
    }

    // [집 번호][고른 색상][시작 색상]
    auto dp = vector<vector<vector<int>>>(
        size+1, vector<vector<int>>(
            3, vector<int>(
                3, 1001
            )
        )
    );
    dp[1][0][0] = costs[1][0];
    dp[1][1][1] = costs[1][1];
    dp[1][2][2] = costs[1][2];

    for (int i=2; i<=size; i++) {
        for (int j=0; j<3; j++) {
            dp[i][0][j] = min(dp[i-1][1][j], dp[i-1][2][j]) + costs[i][0];
            dp[i][1][j] = min(dp[i-1][0][j], dp[i-1][2][j]) + costs[i][1];
            dp[i][2][j] = min(dp[i-1][0][j], dp[i-1][1][j]) + costs[i][2];
        }
    }

    int redMin = min(dp[size][1][0], dp[size][2][0]);
    int greenMin = min(dp[size][0][1], dp[size][2][1]);
    int blueMin = min(dp[size][0][2], dp[size][1][2]);

    int minCost = min(redMin, greenMin);
    minCost = min(minCost, blueMin);

    cout << minCost << '\n';
}