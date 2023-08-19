#include <iostream>

using namespace std;

class Solution {
public:
    int reverse(int x) {
        bool isNegative = (x < 0);
        int reversed = 0;

        while (x != 0) {
            reversed *= 10;
            reversed += x %10;
            
            x /= 10;
        }

        if ((reversed < 0) != isNegative) {
            return 0;
        }

        return reversed;
    }
};