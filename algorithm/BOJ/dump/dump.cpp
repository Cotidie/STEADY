#include <iostream>
#include <vector>
#include <queue>
#include <cmath>

using namespace std;

struct Point {
    float x, y;
};

struct Edge {
    int from, to;
    float cost;
};

class Ascending {
public:
    // right은 새로 들어오는 간선
    bool operator()(const Edge &left, const Edge &right) {
        return left.cost > right.cost;
    }
};

class Disjoint {
public:
    Disjoint(int size) {
        parents = vector<int>(size, -1);
    }

    int find(int n) {
        if (parents[n] < 0) return n;

        int root = find(parents[n]);
        parents[n] = root;

        return root;
    }

    bool unite(int a, int b) {
        int aRoot = find(a);
        int bRoot = find(b);

        if (aRoot == bRoot) return true;

        int aSize = parents[aRoot] * -1;
        int bSize = parents[bRoot] * -1;

        if (aSize > bSize) {
            parents[bRoot] = aRoot;
            parents[aRoot] -= bSize;
        } else {
            parents[aRoot] = bRoot;
            parents[bRoot] -= aSize;
        }

        return false;
    }

private:
    vector<int> parents;
};

class NightSky {
public:
    NightSky(int stars) {
        this->stars = vector<Point>(stars);
        this->edges = vector<Edge>(stars);
    }

    void set(float x, float y, int star) {
        stars[star].x = x;
        stars[star].y = y;
    }

    void connect() {
        for (int i=0; i<stars.size(); i++) {
            for (int j=i+1; j<stars.size(); j++) {
                float xDiff = stars[i].x - stars[j].x;
                float yDiff = stars[i].y - stars[j].y;

                float cost = sqrtf(pow(xDiff, 2) + pow(yDiff, 2));
                edges.push_back({i, j, cost});
            }
        }
    }

    float mst() {
        priority_queue<Edge, vector<Edge>, Ascending> pq(edges.begin(), edges.end());
        Disjoint distSet(stars.size());

        int cnt = 0;
        float cost = 0;
        while (cnt < stars.size() && !pq.empty()) {
            Edge cur = pq.top();
            pq.pop();

            bool isCycle = distSet.unite(cur.from, cur.to);
            if (isCycle) continue;

            cnt++;
            cost += cur.cost;
        }

        return cost;
    }

private:
    vector<Point> stars;
    vector<Edge> edges;
};

int main() {
    ios::sync_with_stdio(false);
    cin.tie(NULL);

    int nums;
    cin >> nums;

    NightSky sky(nums);
    for (int i=0; i<nums; i++) {
        float x, y;
        cin >> x >> y;

        sky.set(x, y, i);
    }
    sky.connect();

    cout << sky.mst() << '\n';
}