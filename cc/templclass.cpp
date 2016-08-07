/*
 * File: templclass.cpp
 * Author: wdeqin
 * Time: 2015/5/24 21:19:06
 * Purpose: c++ template class experiment
 */
#include <iostream>
#include <memory>
#include <cassert>

template <typename T>
class matrix {
private:
    std::unique_ptr<T[]> _data;
    int _row;
    int _col;
public:
    matrix(int row, int col) : _row(row), _col(col) {
        assert(row > 0 && col > 0);
        _data = std::unique_ptr<T[]>(new T[row * col]);
    }

    matrix(int row, int col, T* data) : _row(row), _col(col) {
        assert(row > 0 && col > 0);
        _data = std::unique_ptr<T[]>(new T[row * col]);
        for (int i = 0; i != row * col; ++i) {
            _data[i] = data[i];
        }
    }

    matrix(const matrix& src) : 
        _row(src._row), 
        _col(src._col), 
        _data(new T[src._row * src._col]) {
            T* s = src._data.get();
            T* d = _data.get();
            std::copy(s, s + (src._col * src._res), d);
    }

    matrix& operator=(const matrix& rhs) {
        if (this != &rhs) {
            if (this->_col * this->_row != rhs._col * this->_col) {
                this->_data = nullptr;
            }
            this->_col = rhs._col;
            this->_row = rhs._row;
            this->_data = std::unique_ptr<T[]>(new T[this->_row * this->_col]);
            T* s = rhs._data.get();
            T* d = this->_data.get();
            std::copy(s, s + (this->_row * this->_col), d);
        }
        return *this;
    }

    matrix(matrix&& temp) :
        _row(temp._row),
        _col(temp._col),
        _data(std::move(temp._data)) {
        temp._data = nullptr;
        temp._row = 0;
        temp._col = 0;
    }

    matrix& operator=(matrix&& temp) {
        assert(this != &temp);
        _data = nullptr;
        _data = std::move(temp._data);
        _row = temp._row;
        _col = temp._col;
        temp._data = nullptr;
        temp._row = temp._row;
        temp._col = temp._col;

    }

    ~matrix() {
    }

    friend matrix operator*(const matrix& lhs, const matrix& rhs) {
        assert(lhs._col == rhs._row);
        matrix res(lhs._row, rhs._col);
        for (int r = 0; r != res._row; ++r) {
            for (int c = 0; c != res._col; ++c) {
                int sum = 0;
                for (int k = 0; k != lhs._col; ++k) {
                    sum += lhs._data[r * lhs._col + k] * rhs._data[k * rhs._col + c];
                }
                res._data[r * res._col + c] = sum; 
            }
        }
        return res;
    }

    void show() {
        int tr;
        for (int r = 0; r != _row; ++r) {
            tr = r * _col;
            for (int c = 0; c != _col; ++c) {
                std::cout << _data[tr + c] << "\t";
            }
            std::cout << std::endl;
        }
    }

    void setData(T* data) {
        for (int i = 0; i != _row * _col; ++i) {
            _data[i] = data[i];
        }
    }

    void setData(int row, int col, T value) {
        _data[row * _row + col] = value;
    }

    int getRow() {
        return _row;
    }

    int getCol() {
        return _col;
    }

    const T* getData() {
        return _data;
    }

};

int main() {
    int c[4] = { 1, 1, 1, 1};
    int d[6] = { 1, 2, 3, 4, 5, 6};
    matrix<int> a(2, 2, c);
    std::cout << "matrix a = " << std::endl;
    a.show();
    matrix<int> b(2, 3, d);
    std::cout << "matrix b = " << std::endl;
    b.show();
    matrix<int> r = a * b;
    std::cout << "matrix a * b = " << std::endl;
    r.show();

    return 0;
}
