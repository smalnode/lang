/*
 * File: hello.c
 * Author: wdeqin
 * Time: 2015/5/24 21:17:09
 * Purpose: language c hello world
 */
#include <stdio.h>

// program entry point
int main(int argc, char** argv) {
	if (argc <= 1) {
		printf("Hello, world!\n");
    }
	else {
		printf("Hello, \n");
		int count;
		for (count = 1; count != argc; count++)
			printf("\t%s\n", argv[count]);
	}
	
	return 0;
}
