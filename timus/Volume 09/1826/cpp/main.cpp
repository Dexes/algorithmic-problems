#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;

int main() {
    int n;
    cin >> n;

    vector<int> data(n);
    for (int i = 0; i < n; ++i) {
        cin >> data[i];
    }

    sort(data.begin(), data.end());

    int result = 0;
    while (!data.empty()) {
        if (data.size() == 1) {
            result += data[0];
            data.clear();
            break;
        }

        if (data.size() == 2) {
            result += data[1];
            data.clear();
            break;
        }

        if (data.size() == 3) {
            result += data[0] + data[1] + data[2];
            data.clear();
            break;
        }

        int last = data.back();
        data.pop_back();
        int penultimate = data.back();
        data.pop_back();

        // first, second, penultimate, last
        // penultimate, last -> | <- first, second = second
        // first, penultimate, last -> | <- second = second + first
        // first -> | <- second, penultimate, last = second + first + last
        // first, second -> | <- penultimate, last = second + first + last + second
        int first_case = data[0] + data[1] * 2 + last;

        // first, penultimate, last
        // last -> | <- first, penultimate = penultimate
        // first, last -> | <- penultimate = penultimate + first
        // -> | <- first, penultimate, last = penultimate + first + last
        // first -> | <- penultimate, last = penultimate + first + last + first
        int second_case = data[0] * 2 + penultimate + last;

        result += first_case < second_case ? first_case : second_case;
    }

    cout << result << endl;
    return 0;
}