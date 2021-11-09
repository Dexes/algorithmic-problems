#include <iostream>
#include <stack>

using namespace std;

ostream &operator<<(ostream &stream, stack<char> &n) {
    if (n.empty()) return stream;

    for (auto i = n.size(); i > 0; --i) {
        stream << n.top();
        n.pop();
    }

    return stream;
}

char ReadChar() {
    char ch[1];
    cin.read(ch, 1);
    if (!cin) return 0;

    return *ch;
}

int main() {
    stack<char> n;
    char ch;

    while (ch = ReadChar()) {
        if (isalpha(ch)) {
            n.push(ch);
            continue;
        }

        cout << n << ch;
    }

    cout << n;
    return 0;
}