#include <iostream>

using namespace std;

int main() {
    int a, b;
    cin >> a >> b;

    int sum = a + b - 1;

    cout << (sum - a) << ' ' << (sum - b) << endl;
    return 0;
}