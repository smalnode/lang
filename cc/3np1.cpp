/*
 * File: 3np1.cpp
 * Author: wdeqin
 * Time: 2015/5/24 21:15:51
 * Purpose: the 3n + 1 and n/2 loop
 */
#include <iostream>
using namespace std;

int a[100];

int maxCircle(int i, int j);
int circle(int i);

int main() {
    int i, j, k;
    k = 0;
    while (cin >> i && cin >> j && !cin.eof()) {
        a[k++] = i;
        a[k++] = j;
    }

    for (int t = 0; t < k; t += 2) {
        cout << a[t] << " " << a[t + 1] << " "
            << maxCircle(a[t], a[t + 1]) << endl;
    }
    return 0;
}

int maxCircle(int i, int j) {
    int t, s;
    t = circle(i);
    for (int k = i + 1; k <= j; k++) {
        s = circle(k);
        if (s > t)
            t = s;
    }
    return t;
}

int circle(int i) {
    int c = 1;
    while (i != 1) {
        i = i & 1 ? i * 3 + 1 : i / 2;
        c++;
    }
    return c;
}
