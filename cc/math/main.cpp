#include <iostream>
#include "dcomplex.h"
using namespace std;
using namespace dmath;

int main() {
	dcomplex a(1, 2);
	dcomplex b(a);

	dcomplex c = a * b;

	cout << "a = " << a << endl;
	cout << "b = " << b << endl;
	cout << "c = a * b / 2 = " << c / 2 << endl;
    cout << "d = c - a = " << c - a << endl;

	return 0;
}
