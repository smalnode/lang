/*
 * File: nth_element.c
 * Author: wdeqin
 * Time: Sun 04 Sep 2016 07:11:09 PM CST
 * Purpose: Use quicksort find the nth biggest element in unsorted array
 */

#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

int nth_element(int* arr, int bgn, int end, int n);

int main() {
    int len;
    printf("input num of element:\n");
    scanf("%d", &len);
    int* arr = (int*)malloc(sizeof(int) * len);

    printf("input %d integer:\n", len);
    int i = 0;
    while (i < len) {
        scanf("%d", arr + i);
        i++;
    }
    int n;
    printf("input n th:\n");
    scanf("%d", &n);
    printf("%dth element: %d\n", n, nth_element(arr, 0, len - 1, n));
    free(arr);
    arr = 0x0;
    return 0;
}

int nth_element(int* arr, int bgn, int end, int n) {
    if (bgn >= end) {
        return arr[end];
    }
    int p = arr[bgn];
    int i = bgn;
    int j = bgn + 1;
    int tmp;
    while (j <= end) {
        if (arr[j] < p) {
            // swap arr[i + 1] and arr[j]
            tmp = arr[i + 1];
            arr[i + 1] = arr[j];
            arr[j] = tmp;
            i++;
        }
        j++;
    }
    tmp = arr[bgn];
    arr[bgn] = arr[i];
    arr[i] = tmp;

    // [bgn ~ i- i i+ end]
    int gep = end - i + 1;
    if (n <= gep) {
        return nth_element(arr, i + 1, end, n);
    } else {
        return nth_element(arr, bgn, i - 1, n - gep);
    }
}
