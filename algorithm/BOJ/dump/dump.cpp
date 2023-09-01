#include <iostream>
#include <vector>
#include <queue>

using namespace std;

const int INTMAX = 200000000;

struct Pair {
    int node;
    int cost;
};

struct Edge {
    int to;
    int cost;
};

class Ascending {
public:
    bool operator()(const Pair &left, const Pair &right) {
        return left.cost > right.cost;
    }
};

class Graph {
public:
    Graph(int size) {
        graph = vector<vector<Edge>>(size, vector<Edge>());
    }

    void set(int from, int to, int cost) {
        graph[from].push_back({to, cost});
        graph[to].push_back({from, cost});
    }

    vector<int> distances(int start) {
        vector<int> dist = vector<int>(graph.size(), -1);
        dist[start] = 0;
        
        priority_queue<Pair, vector<Pair>, Ascending> pq;
        pq.push({start, 0});

        while (!pq.empty()) {
            Pair cur = pq.top();
            pq.pop();

            if (dist[cur.node] != -1 && cur.cost > dist[cur.node]) continue;

            for (auto edge : graph[cur.node]) {
                int curDist = dist[edge.to];
                int newDist = cur.cost + edge.cost;
                if (curDist != -1 && curDist <= newDist) continue;

                dist[edge.to] = newDist;
                pq.push({edge.to, newDist});
            }
        }

        return dist;
    }

private:
    vector<vector<Edge>> graph;
};

int main() {
    int nodes, edges;
    cin >> nodes >> edges;

    Graph g(nodes);
    for (int i=0; i<edges; i++) {
        int from, to, cost;
        cin >> from >> to >> cost;

        g.set(from-1, to-1, cost);
    }

    int med1, med2;
    cin >> med1 >> med2;
    med1--; med2--;

    vector<int> distStart = g.distances(0);
    vector<int> distMed1 = g.distances(med1);
    vector<int> distMed2 = g.distances(med2);

    int shortest = -1;

    int startToMed1 = distStart[med1];
    int medToMed = distMed1[med2];
    int med1ToEnd = distMed1[nodes-1];
    int med2ToEnd = distMed2[nodes-1];

    if (medToMed == -1 || med1ToEnd == -1 || med2ToEnd == -1) {
        cout << -1 << '\n';
        return 0;
    }

    shortest = min(
        // 시작 -> i -> j -> 도착
        distStart[med1] + distMed1[med2] + distMed2[nodes-1],
        // 시작 -> j -> i -> 도착
        distStart[med2] + distMed2[med1] + distMed1[nodes-1]
    );
    
    cout << shortest << '\n';
}