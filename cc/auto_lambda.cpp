/*
 * File: auto_lambda.cpp
 * Author: wdeqin
 * Time: Mon 31 Aug 2015 12:37:46 AM CST
 * Purpose: Test c++14 auto parm of lambda function
 */
#include <iostream>
using namespace std;

typedef int (*intFilter) (int parm);

int filtInt(intFilter f, int parm) {
    return f(parm);
}

typedef double (*doubleFilter) (double parm);

double filtDouble(doubleFilter f, double parm) {
    return f(parm);
}

int main() {
    // so that anymous function a actiong like generic function with auto term
    auto a = [](auto parm) { return parm * parm; };
    cout << filtInt(a, 10) << endl;
    cout << filtDouble(a, 10) << endl;
    return 0;
}
