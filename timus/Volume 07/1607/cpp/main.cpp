#include<iostream>

using namespace std;

int main() {
    int a, b, c, d;
    cin >> a >> b >> c >> d;

    if (a > c) {
        cout << a << endl;
        return 0;
    }

    while (true) {
        if (a + b <= c) {
            a += b;
        } else {
            cout << c << endl;
            return 0;
        }

        if (c - d >= a) {
            c -= d;
        } else {
            cout << a << endl;
            return 0;
        }
    }
}
