/*
 * File: eof.c
 * Author: wdeqin
 * Time: 2015/5/24 9:05:19
 * Purpose: eof test
 */
#include <stdio.h>

/* Program entry point */
int main(void) {
	char c;
	c = getchar();
	while(c != EOF) {
		printf("%c", c);
		c = getchar();
	}

	return 0;
} /* end main */
