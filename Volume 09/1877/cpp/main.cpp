#include <iostream>

using namespace std;

int main() {
    int first, second;
    cin >> first >> second;

    cout << (first % 2 == 0 || second % 2 != 0 ? "yes" : "no") << endl;
    return 0;
}