/*
 * File: derive.cpp
 * Author: wdeqin
 * Time: 2015/5/24 21:16:51
 * Purpose: c++ inheritance, virtual function
 */
#include <iostream>
using std::cout;
using std::endl;

class B {
public:
	virtual void foo() {
		cout << "B::foo()" << endl;
	}
};


class Ba : public B {
public:
	virtual void foo() {
		cout << "Ba::foo()" << endl;
	}
};

class Bc : public B {
public:
	virtual void foo() override {
		cout << "Bc::foo()" << endl;
	}
};

int main() {
	B* b = new B();
	b->foo();

	delete b;
	b = nullptr;

	b = new Ba();
	b->foo();

	Ba* ba = (Ba*) b;
	ba->foo();

	delete b;
	b = nullptr;

	b = new Bc();
	b->foo();

	delete b;
	b = nullptr;
}
