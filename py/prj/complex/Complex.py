__name__ = 'complex'

class Complex():
    def __init__(self, re, im):
        self._re = re
        self._im = im

    def __repr__(self):
        return "Complex({}, {})".format(self._re, self._im)

def compmulti(lhs, rhs):
    if not isinstance(lhs, Complex) or not isinstance(rhs, Complex):
        raise TypeError('Not Complex tyoe')
    return Complex(lhs._re * rhs._re - lhs._im * rhs._im,
            lhs._re * rhs._im + lhs._im * rhs._re)
