#include <iostream>
#include <vector>
#include <queue>

using namespace std;

class Topsort {
public:
    Topsort(int size) {

    }

    void sort() {

    }
private:
    vector<int> nums;
};

int main() {
    int nums, deps;
    cin >> nums >> deps;

    Topsort topsort(nums);
    for (int i=0; i<deps; i++) {
        int first, after;
        cin >> first >> after;

        solver.set(after, first);
    }

    solver.solve();
    cout << '\n';
}