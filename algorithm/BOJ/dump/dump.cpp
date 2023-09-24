#include <iostream>
#include <vector>
#include <string>
#include <map>

using namespace std;

struct Node {
    string name = "";
    int size = 1;
    Node *parent = nullptr;
};

class Disjoint {
public:
    Disjoint() {}

    Node* find(string name) {
        if (!exists(name)) {
            parents[name] = {name, 1, nullptr};
            return &parents[name];
        }

        Node *cur = &parents[name];
        if (cur->parent == nullptr) return cur;

        Node *root = find(cur->parent->name);
        cur->parent = root;

        return root;
    }

    bool unite(string a, string b) {
        Node *aRoot = find(a);
        Node *bRoot = find(b);

        if (aRoot == bRoot) return true;

        if (aRoot->size < bRoot->size) {
            aRoot->parent = bRoot;
            bRoot->size += aRoot->size; 
        } else {
            bRoot->parent = aRoot;
            aRoot->size += bRoot->size;
        }

        return false;
    }

    int getSize(string name) {
        return find(name)->size;
    }

private:
    bool exists(string name) {
        return parents.find(name) != parents.end();
    }

    map<string, Node> parents;
};

int main() {
    ios::sync_with_stdio(false);
    cin.tie(NULL);

    int T;
    cin >> T;

    for (int t=0; t<T; t++) {
        int nums;
        cin >> nums;

        Disjoint distSet;
        for (int n=0; n<nums; n++) {
            string a, b;
            cin >> a >> b;

            distSet.unite(a, b);
            cout << distSet.getSize(a) << '\n';
        }
    }
}