#include <iostream>
#include <vector>

using namespace std;

struct Point {
    int left=0;
    int right=0;
};

int getCost(int from, int to) {
    if (from == 0) return 2;
    if (abs(from-to) == 2) return 4;
    if (from == to) return 1;

    return 3; 
}

int main() {
    vector<int> inputs;
    
    int input;
    while (cin >> input, input != 0) {
        inputs.push_back(input);
    }
    
    int size = inputs.size();
    
    vector<vector<int>> dp(size+1, vector<int>(2, 0));
    vector<vector<Point>> location(size+1, vector<Point>(2));

    int minCost;
    Point minPoint;
    for (int i=1; i<=size; i++) {
        int input = inputs[i-1];
        
        int leftMoveFromLeft = dp[i-1][0] + getCost(location[i-1][0].left, input);
        int leftMoveFromRight = dp[i-1][1] + getCost(location[i-1][1].left, input);
        if (leftMoveFromLeft <= leftMoveFromRight) {
            dp[i][0] = leftMoveFromLeft;
            location[i][0] = Point{input, location[i-1][0].right};
        } else {
            dp[i][0] = leftMoveFromRight;
            location[i][0] = Point{input, location[i-1][1].right};
        }

        int rightMoveFromLeft = dp[i-1][0] + getCost(location[i-1][0].right, input);
        int rightMoveFromRight = dp[i-1][1] + getCost(location[i-1][1].right, input);
        if (rightMoveFromLeft <= rightMoveFromRight) {
            dp[i][1] = rightMoveFromLeft;
            location[i][1] = Point{location[i-1][0].left, input};
        } else {
            dp[i][1] = rightMoveFromRight;
            location[i][1] = Point{location[i-1][1].left, input};
        }
    }

    cout << min(dp[size][0], dp[size][1]) << '\n';
}