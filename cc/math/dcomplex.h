//~

#include <ostream>

#ifndef __COMPLEX_H__
#define __COMPLEX_H__

namespace dmath {

class dcomplex {
	public:
		double re;
		double im;

		dcomplex();
		dcomplex(double re, double im);
		dcomplex(const dcomplex& v);
		double abs() const;
};

const dcomplex operator+(const dcomplex&, const dcomplex&);
const dcomplex operator-(const dcomplex&, const dcomplex&);
const dcomplex operator*(const dcomplex&, const dcomplex&);
const dcomplex operator*(const dcomplex&, const double&);
const dcomplex operator*(const dcomplex&, const int&);
const dcomplex operator/(const dcomplex&, const double&);
double abs(const dcomplex&);
std::ostream& operator<<(std::ostream&, const dcomplex&);

}

#endif // __COMPLEX_H__
