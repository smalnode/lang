/*
 * File: hello.cpp
 * Author: wdeqin
 * Time: 2015/5/24 21:17:41
 * Purpose: language c++ hello world
 */
#include <iostream>

int main(int argc, char* argv[]) {
    for (int i = 0; i != argc; ++i) {
        std::cout << "|" << argv[i] << "|\n";
    }
    
    int i  = 0;
    int j = 0;
    int b = i++;
    int c = ++j;

    std::cout << "0++ = " << b << std::endl;
    std::cout << "++0 = " << c << std::endl;  

    return 0;
}
