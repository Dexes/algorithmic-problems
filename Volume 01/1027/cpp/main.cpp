#include <iostream>
#include <algorithm>
#include <array>
#include <queue>

using namespace std;

bool is_comment = false;
bool is_math = false;
int bracket_balance = 0;

queue<char> input;
char current_char, next_char;
array<char, 15> math_chars = {'=', '+', '-', '*', '/', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'};

char pop_char() {
    if (!input.empty()) {
        char result = input.front();
        input.pop();

        return result;
    }

    char result;
    while (true) {
        result = getc(stdin);
        if (result != 10 && result != 13) {
            return result;
        }
    }
}

void push_char(char ch) {
    input.push(ch);
}

bool check_text() {
    for (current_char = pop_char(); current_char != EOF; current_char = pop_char()) {
        // check comment
        if (!is_comment && current_char == '(') {
            next_char = pop_char();
            if (next_char == '*') {
                is_comment = true;
                continue;
            }

            push_char(next_char);
        }

        if (is_comment && current_char == '*') {
            next_char = pop_char();
            if (next_char == ')') {
                is_comment = false;
                continue;
            }

            push_char(next_char);
        }

        if (is_comment) {
            continue;
        }

        // check math
        if (current_char == '(') {
            is_math = true;
            ++bracket_balance;
            continue;
        }

        if (is_math && current_char == ')') {
            is_math = --bracket_balance != 0;
            continue;
        }

        if (is_math) {
            bool in_array = find(math_chars.begin(), math_chars.end(), current_char) != math_chars.end();
            if (!in_array) {
                return false;
            }

            continue;
        }

        // check text
        if (current_char == '(' || current_char == ')') {
            return false;
        }
    }

    return !is_comment && !is_math;
}

int main() {
    cout << (check_text() ? "YES" : "NO") << endl;
    return 0;
}