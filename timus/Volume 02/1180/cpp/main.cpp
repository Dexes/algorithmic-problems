#include <stdio.h>

int main() {
    char ch;
    int sumall = 0;
    while (!feof(stdin)) {
        scanf("%c", &ch);
        if (ch < '0' || ch > '9') {
            break;
        }

        sumall += ch;
    }

    sumall %= 3;
    if (sumall) {
        printf("1\n%d", sumall);
    } else {
        printf("2");
    }

    return 0;
}
