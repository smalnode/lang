/*
 * File: static_assert.cpp
 * Author: wdeqin
 * Time: 2015/5/24 21:18:51
 * Purpose: c++11 static_assert
 */
#include <type_traits>

int main() {
	static_assert(1 > 2, "1 > 2");

	return 0;
}
