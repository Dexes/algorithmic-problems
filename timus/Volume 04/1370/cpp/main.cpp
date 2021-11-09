#include <iostream>
#include <vector>

using namespace std;

int main() {
    int count, clicks;
    cin >> count >> clicks;

    vector<char> data(count);
    for (int i = 0; i < count; ++i) {
        cin >> data[i];
    }

    for (int i = clicks; i < clicks + 10; ++i) {
        cout << data[i % count];
    }

    cout << endl;
    return 0;
}