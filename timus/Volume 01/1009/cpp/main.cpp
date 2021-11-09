#include <iostream>

using namespace std;

int main() {
    int n, k;
    cin >> n >> k;

    int tmp, first = k - 1, second = first * k, current = 2;

    while (current != n) {
        tmp = (first + second) * (k - 1);
        first = second;
        second = tmp;
        ++current;
    }

    cout << second << endl;
    return 0;
}