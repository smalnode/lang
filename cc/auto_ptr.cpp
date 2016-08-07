/*
 * File: auto_ptr.cpp
 * Author: wdeqin
 * Time: 2015/5/24 22:58:17
 * Purpose: c++11 auto pointer
 */
#include <iostream>

int main() {
    int v = 17;
    auto* p = &v;
    std::cout << "*(auto* p = &(int v = 17)) = " << *p << std::endl;
    return 0;
}
