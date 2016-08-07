//~

#include "dcomplex.h"
#include <math.h>
#include <iostream>
#include <math.h>

namespace dmath {

dcomplex::dcomplex() : re(.0), im(.0) { }

dcomplex::dcomplex(double re, double im) : re(re), im(im) { }

dcomplex::dcomplex(const dcomplex& v) : re(v.re), im(v.im) { }

double dcomplex::abs() const {
	return sqrt(this->re * this->re + this->im * this->im);
}

const dcomplex operator+(const dcomplex& lhs, const dcomplex& rhs) {
	return dcomplex(lhs.re + rhs.re, lhs.im + rhs.im);
}

const dcomplex operator-(const dcomplex& lhs, const dcomplex& rhs) {
	return dcomplex(lhs.re - rhs.re, lhs.im - rhs.im);
}

const dcomplex operator*(const dcomplex& lhs, const dcomplex& rhs) {
	return dcomplex(lhs.re * rhs.re - lhs.im * rhs.im,
			lhs.re * rhs.im + lhs.im * rhs.re);
}

const dcomplex operator*(const dcomplex& lhs, const double& rhs) {
	return dcomplex(lhs.re * rhs, lhs.im * rhs);
}

const dcomplex operator*(const dcomplex& lhs, const int& rhs) {
	return dcomplex(lhs.re * rhs, lhs.im * rhs);
}

const dcomplex operator/(const dcomplex& lhs, const double& rhs) {
	return dcomplex(lhs.re / rhs, lhs.im / rhs);
}

std::ostream& operator<<(std::ostream& output, 
		const dcomplex& c) {
	output << c.re << "+" << c.im << "i";
	return output;
}

double abs(const dcomplex& v) {
	return v.abs();
}

}
