/*
 * File: const.cpp
 * Author: wdeqin
 * Time: 2015/5/24 21:16:21
 * Purpose: Const ptr point to const int
 */
#include <iostream>
using namespace std;

int main() {
	const int a = 100;
	const int b = 50;
	const int* const p = &a; 
    // first const indicate p point to a const, second indicate p is a const it self
	cout << *p << endl;
	p = &a;
	return 0;
}
