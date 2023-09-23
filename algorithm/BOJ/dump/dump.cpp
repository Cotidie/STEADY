#include <iostream>
#include <vector>

using namespace std;

struct Point {
    int row, col;
};

class Disjoint {
public:
    Disjoint(int rows, int cols) {
        parents = vector<vector<Point>>(
            rows, vector<Point>(cols, Point{-1, -1})
        );
    }

    // 참조자로 개선 가능
    Point find(Point p) {
        if (isRoot(p)) return p;

        Point root = find(parents[p.row][p.col]);
        parents[p.row][p.col] = root;

        return root;
    }

    bool unite(Point a, Point b) {
        Point aRoot = find(a);
        Point bRoot = find(b);

        if (isSamePoint(aRoot, bRoot)) return true;

        parents[a.row][a.col] = bRoot;
        return false;
    }

private:
    bool isRoot(Point &p) {
        return parents[p.row][p.col].row == -1;
    }

    bool isSamePoint(Point &a, Point &b) {
        return a.row == b.row && a.col == b.col;
    }

    vector<vector<Point>> parents;
};

class Board {
public:
    Board(int rows, int cols): rows(rows), cols(cols) {
        board = vector<string>(rows);
    }

    void set(int row, string cols) {
        board[row] = cols;
    }

    int makeSafezones() {
        Disjoint distSet(rows, cols);

        int cnt = 0;
        for (int r=0; r<rows; r++) {
            for (int c=0; c<cols; c++) {
                Point cur = {r, c};
                Point next = getNext(r, c);
                if (!isInbound(next.row, next.col)) {
                    cnt++;
                    continue;
                }

                bool isCycle = distSet.unite(cur, next);
                if (isCycle) cnt++;
            }
        }

        return cnt;
    }

private:
    Point getNext(int row, int col) {
        char dir = board[row].at(col);

        if (dir == 'L') return {row, col-1};
        if (dir == 'R') return {row, col+1};
        if (dir == 'U') return {row-1, col};
        if (dir == 'D') return {row+1, col};

        return {row, col};
    }

    bool isInbound(int row, int col) {
        return (row >= 0 && col >= 0) &&
            (row < rows && col < cols);
    }

    vector<string> board;
    int rows, cols;
};

int main() {
    int rows, cols;
    cin >> rows >> cols;

    Board board(rows, cols);
    for (int i=0; i<rows; i++) {
        string val;
        cin >> val;

        board.set(i, val);
    }

    cout << board.makeSafezones() << '\n';
}