#include <iostream>
#include <vector>

using namespace std;

struct Point {
    int row, col;
};

class City {
public:
    City() {};

    void setHouse(int row, int col) {
        houses.push_back({row, col});
    }

    void setShop(int row, int col) {
        shops.push_back({row, col});
    }

    int openShops() {
        int min = 10000000;

        for (auto chickens : combinations) {
            int sum = 0;

            for (auto house : houses) {
                int dist = 10000000;
                for (auto chicken : chickens) {
                    int newDist = getDist(house, chicken);
                    if (dist > newDist) dist = newDist;
                }

                sum += dist;
            }

            if (min > sum) min = sum;
        }

        return min;
    }

    void findCombinations(int size) {
        auto t = vector<Point>(size);
        _findCombinations(t, 0, 0, size);
    }

private:
    int getDist(Point &p1, Point &p2) {
        return abs(p1.row - p2.row) + abs(p1.col - p2.col);
    }

    void _findCombinations(vector<Point> &comb, int idx, int depth, int size) {
        if (depth == size) {
            combinations.push_back(comb);
            return;
        }

        for (int i=idx; i<shops.size(); i++) {
            comb[depth] = shops[i];
            _findCombinations(comb, i+1, depth+1, size);
        }
    }

    vector<Point> houses;
    vector<Point> shops;

    vector<vector<Point>> combinations;
};

int main() {
    int size, shops;
    cin >> size >> shops;

    City city;
    for (int r=0; r<size; r++) {
        for (int c=0; c<size; c++) {
            int val;
            cin >> val;

            if (val == 1) city.setHouse(r, c);
            if (val == 2) city.setShop(r, c);
        }
    }

    city.findCombinations(shops);
    
    cout << city.openShops() << '\n';
}