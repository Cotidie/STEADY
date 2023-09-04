#include <iostream>
#include <cmath>

using namespace std;

float roundUnder(float num, int digit) {
    float div = pow(10, digit);

    return round(num*div) / div;
}

int main() {
    double weight = 0.995;
    
    double score = 0;
    for (int i=1; i<400; i++) {
        double curScore = 175 * (1- pow(weight, i));
        cout << i << ". " << 175 * (1- pow(weight, i)) << ", " << curScore - score << '\n';
        score = curScore;
    }
}