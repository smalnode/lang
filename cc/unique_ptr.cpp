/*
 * File: unique_ptr.cpp
 * Author: wdeqin
 * Time: 2015/5/24 21:19:23
 * Purpose: c++11 unique_ptr
 */
#include <memory>
#include <iostream>

int main() {
    std::unique_ptr<double> a(new double);
    std::cout << "before move" << std::endl;
    std::cout << "a = " << a.get() << std::endl;
    std::cout << "*a = " << *a.get() << std::endl;
    std::unique_ptr<double> c = std::move(a);
    std::cout << "after move" << std::endl;
    std::cout << "a = " << a.get() << std::endl;
    std::cout << "c = " << c.get() << std::endl;
    std::cout << "*c = " << *c.get() << std::endl;
    return 0;
}


