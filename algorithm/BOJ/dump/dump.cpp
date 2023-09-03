#include <iostream>
#include <vector>

using namespace std;

struct Point {
    int row, col;
};

class Room {
public:
    Room(int rows, int cols):rows(rows), cols(cols) {
        room = vector<vector<int>>(rows, vector<int>(cols, 0));
    }

    int get(int row, int col) { return room[row][col]; }

    void set(int row, int col, int val) {
        if (val == -1) conditioner.set(row, col);
        if (val > 0) dirts += val;
        
        room[row][col] = val;
    }

    void next() {
        vector<vector<int>> spread = vector<vector<int>>(rows, vector<int>(cols, 0)); 
        calculateSpread(spread);
        applySpread(spread);

        dirts -= conditioner.push(*this, "up");
        dirts -= conditioner.push(*this, "down");
    }

    int get() { return dirts; }
    int getCols() {return cols;}
    int getRows(){return rows;}

    bool isOutOfRange(int row, int col) {
        return (row >= 0 && col >= 0) && 
            (row < rows && col < cols);
    }

private:
    void calculateSpread(vector<vector<int>> &spread) {
        for (int r=0; r<rows; r++) {
            for (int c=0; c<cols; c++) {
                if (room[r][c] == -1 || room[r][c] == 0) continue;
                
                spreadCell(r, c, spread);
            }
        }
    }

    void spreadCell(int row, int col, vector<vector<int>> &spread) {
        int amount = room[row][col] / 5;
        int cnt = 0;

        for (Point &move : moves) {
            int drow = row+move.row;
            int dcol = col+move.col;

            if (isOutOfRange(drow, dcol)) continue;
            if (room[drow][dcol] == -1) continue;

            spread[drow][dcol] += amount;
            cnt++;
        }

        spread[row][col] -= cnt * amount;
    }

    void applySpread(vector<vector<int>> &spread) {
        for (int r=0; r<rows; r++) {
            for (int c=0; c<cols; c++) {
                room[r][c] += spread[r][c];
            }
        }
    }

    int rows, cols;
    int dirts = 0;
    
    vector<vector<int>> room;
    vector<Point> moves = {
        {1, 0}, {-1, 0}, {0, 1}, {0, -1}
    };
    Conditioner conditioner;
};

class Conditioner {
public:
    Conditioner(Room &room):room(room) {};
    void set(int row, int col);
    int push(string dir);
    int _push(vector<vector<int>> &dirs);

private:
    bool isUpperSet() { return up.row != -1; }

    Room& room;
    vector<vector<int>> dirUp = {
        {0, 1}, {-1, 0}, {0, -1}, {1, 0}
    };
    vector<vector<int>> dirDown = {
        {0, 1}, {1, 0}, {0, -1}, {1, 0}
    };

    Point up = {-1, -1};
    Point down = {-1, -1};
};

void Conditioner::set(int row, int col) {
        if (isUpperSet()) down = {row, col};
        else up = {row, col};
    }

int Conditioner::push(Room &room, string dir) {
    if (dir == "up") return _push(room, dirUp);
    else return _push(room, dirDown);
}

int Conditioner::_push(Room &room, vector<vector<int>> &dirs) {
    int absorbed = 0;

    int row = up.row, col = up.col+1;
    int prev = 0, temp = 0;

    int dir = 0;
    while (temp = room.get(row, col), temp != -1) {
        room.set(row, col, prev);
        prev = temp;

        int dRow = row + dirs[dir][0];
        int dCol = col + dirs[dir][1];
        if (room.isOutOfRange(dRow, dCol)) {
            dir++;
            dRow = row + dirs[dir][0];
            dCol = col + dirs[dir][1];
        }

        row = dRow; col = dCol;
    }

    absorbed = prev;
    return absorbed;
}

int main() {
    int rows, cols, seconds;
    cin >> rows >> cols >> seconds;

    Room room(rows, cols);
    for (int r=0; r<rows; r++) {
        for (int c=0; c<cols; c++) {
            int val;
            cin >> val;

            room.set(r, c, val);
        }
    }

    for (int t=0; t<seconds; t++) {
        room.next();
    }

    cout << room.get() << '\n';
}