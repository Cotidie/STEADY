#include <iostream>
#include <string>

using namespace std;

int replaceAll(string &str, string before, string after) {
    int cnt=0;

    int pos=0;
    while (pos = str.find(before), pos != string::npos) {
        str.replace(pos, before.length(), after);
        pos += after.length();
        cnt++;
    }

    return cnt;
}

int main() {
    string str;
    cin >> str;

    string bomb;
    cin >> bomb;

    int cnt = replaceAll(str, bomb, "");

    if (str == "") cout << "FRULA" << '\n';
    else cout << str << '\n';
}