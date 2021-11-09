#include<iostream>

#define MAX 101

int main() {
    int table[MAX][MAX];
    int n, i, j, k, t, max;
    bool all_is_negative = true;

    scanf("%d", &n);
    memset(table, 0, sizeof(table));

    max = -127;
    for (i = 1; i <= n; i++) {
        for (j = 1; j <= n; j++) {
            scanf("%d", &table[i][j]);
            all_is_negative = all_is_negative && (table[i][j]) < 0;
            if (table[i][j] > max) {
                max = table[i][j];
            }
        }
    }

    if (all_is_negative) {
        printf("%d\n", max);
        return 0;
    }

    for (j = 1; j <= n; j++) {
        for (i = 1; i <= n; i++) {
            table[i][j] = table[i - 1][j] + table[i][j];
        }
    }

    for (i = 1; i <= n; i++) {
        for (j = i; j <= n; j++) {
            t = 0;
            for (k = 1; k <= n; k++) {
                t += table[j][k] - table[i - 1][k];
                if (t < 0) t = 0;
                if (t > max) max = t;
            }
        }
    }

    printf("%d\n", max);
    return 0;
}
