/*
 * File: auto.cpp
 * Author: wdeqin
 * Time: 2015/5/22 17:25:38
 * Purpose: c++11 auto experiment
 */
#include <iostream>
using namespace std;

class tim {
public:
	tim(double t) : t(t) {}
	double t;
};

class speed {
public:
	speed(double s) : s(s) {}
	double s;
};

class dis {
public:
	dis(double d) : d(d) {}
	double d;
};

const dis operator*(const speed& s, const tim& t) {
	return dis(s.s * t.t);
}

const dis operator*(const tim& t, const speed& s) {
	return dis(s.s * t.t);
}

ostream& operator<<(ostream& output, const dis& d) {
	output << d.d;
	return output;
}

template <typename T1, typename T2>
auto multiply(T1 t1, T2 t2) -> decltype(t1 * t2) {
	return t1 * t2;
}

int main() {
	speed s(3.1415926);
	tim t(10);

	auto d = multiply(t, s);

	cout << d << endl;

	return 0;
}
