#include <iostream>
#include <map>

using namespace std;

int main(int argc, char **argv)
{
    map<string, int> set;
    string s, sub;

    cin >> s;

    for (int size = 1; size <= s.length(); size++)
    {
        for (int i = 0; i <= (s.length() - size); i++)
        {
            set[s.substr(i, size)] = 1;
        }
    }

    cout << set.size() << '\n';
}