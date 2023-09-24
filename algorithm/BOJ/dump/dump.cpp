#include <iostream>
#include <vector>

long INTMAX = 100001;

using namespace std;

class Floyd {
public:
    Floyd(int size): size(size+1) {
        cost = vector<vector<long>>(size+1, vector<long>(size+1, INTMAX));
        path = vector<vector<int>>(size+1, vector<int>(size+1, -1));
    }

    void set(int city1, int city2, int val) {
        if (val > cost[city1][city2]) return;

        cost[city1][city2] = val;
        path[city1][city2] = city2;
        
    }

    void findShortestPath() {
        for (int k=1; k<size; k++) {
            for (int r=1; r<size; r++) {
                if (k == r) continue;
                for (int c=1; c<size; c++) {
                    if (r==c) continue;

                    long med1 = cost[r][k];
                    long med2 = cost[k][c];
                    if (med1+med2 < cost[r][c]) {
                        cost[r][c] = med1+med2;
                        path[r][c] = path[r][k];
                    }
                }
            }
        }
    }

    void printCosts() {
        for (int r=1; r<size; r++) {
            for (int c=1; c<size; c++) {
                if (cost[r][c] == INTMAX) cout << 0 << " ";
                else cout << cost[r][c] << " ";
            }
            cout <<'\n';
        }
    }

    void printPaths(int start) {
        for (int c=1; c<size; c++) {
            if (!isPathExists(start, c)) {
                cout << 0 << '\n';
                continue;
            }

            vector<int> trace;
            getTrace(start, c, trace);

            cout << trace.size() << " ";
            for (int i=0; i<trace.size(); i++) {
                cout << trace[i] << " ";
            }
            cout << '\n';
        }
    }
private:
    bool isPathExists(int start, int dest) {
        return cost[start][dest] != INTMAX;
    }

    void getTrace(int start, int dest, vector<int> &trace) {
        if (start == dest) {
            trace.push_back(dest);
            return;
        }

        trace.push_back(start);
        start = path[start][dest];

        getTrace(start, dest, trace);
    }

    vector<vector<long>> cost;
    // path[start][end]: start 다음에 방문할 노드
    vector<vector<int>> path;
    int size;
};

int main() {
    int cities, buses;
    cin >> cities >> buses;

    Floyd floyd(cities);
    for (int b=0; b<buses; b++) {
        int start, dest, cost;
        cin >> start >> dest >> cost;

        floyd.set(start, dest, cost);
    }

    floyd.findShortestPath();

    floyd.printCosts();
    for (int i=1; i<=cities; i++) {
        floyd.printPaths(i);
    }
}