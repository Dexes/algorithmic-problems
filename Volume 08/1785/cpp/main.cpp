#include <iostream>

using namespace std;

int main() {
    int count;
    cin >> count;

    if (count < 5) {
        cout << "few" << endl;
        return 0;
    }

    if (count < 10) {
        cout << "several" << endl;
        return 0;
    }

    if (count < 20) {
        cout << "pack" << endl;
        return 0;
    }

    if (count < 50) {
        cout << "lots" << endl;
        return 0;
    }

    if (count < 100) {
        cout << "horde" << endl;
        return 0;
    }

    if (count < 250) {
        cout << "throng" << endl;
        return 0;
    }

    if (count < 500) {
        cout << "swarm" << endl;
        return 0;
    }

    if (count < 1000) {
        cout << "zounds" << endl;
        return 0;
    }

    cout << "legion" << endl;
    return 0;
}