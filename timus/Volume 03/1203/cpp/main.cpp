#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;

struct Section {
    int timeStart;
    int timeEnd;
};

int main() {
    int n, start, end;
    cin >> n;

    vector<Section> sections(n);
    for (int i = 0; i < n; ++i) {
        cin >> start >> end;
        sections[i] = {start, end};
    }

    sort(sections.begin(), sections.end(), [](const Section &a, const Section &b) -> bool {
        return a.timeEnd < b.timeEnd;
    });

    int count = 1, last = 0;
    for (int i = 0; i < n; ++i)
        if (sections[i].timeStart > sections[last].timeEnd) {
            ++count;
            last = i;
        }

    cout << count << endl;
    return 0;
}