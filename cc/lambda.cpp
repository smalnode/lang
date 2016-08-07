/*
 * File: lamda.cpp
 * Author: wdeqin
 * Time: 2015/5/24 21:18:02
 * Purpose: c++11 lambda experiment
 */
#include <iostream>
#include <vector>
#include <algorithm>
#include <functional>
using namespace std;

int main() {
    vector<int> v;
    v.push_back(1);
    v.push_back(2);
    v.push_back(4);
    v.push_back(8);
    v.push_back(16);

    for_each(v.begin(), v.end(), [](int n) { cout << n << endl; });

    auto pos = find_if(v.begin(), v.end(), [](int n) { return (n & 1) == 0; });

    if (pos != v.end()) {
        cout << *pos << endl;
    }

    std::function<int(int)> fib =
        [&fib](int n) { return n < 2 ? 1 : fib(n - 1) + fib(n - 2); };

    cout << fib(10) << endl;

    []() { cout << "this is a anymous function" << endl; }();

    return 0;
}
