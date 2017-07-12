#include <stdio.h>

void checkint(int got, int want) {
    if (got != want) {
        printf("FAILED %d != %d\n", got, want);
    } else {
        printf("passed\n");
    }
}

void printarray(int *array, int length) {
    for (int i = 0; i < length; ++i) {
        printf("%d", array[i]);
        if (i != length-1) {
            printf(" ");
        }
    }
}

void checkarray(int *array1, int *array2, int length) {
    int areEqual = 1;
    for (int i = 0; i < length; ++i) {
        if (array1[i] != array2[i]) {
            areEqual = array1[i] == array2[i];
            break;
        }
    }
    if (areEqual) {
        printf("passed\n");
    } else {
        printf("FAILED [");
        printarray(array1, length);
        printf("] != [");
        printarray(array2, length);
        printf("]\n");
    }
}
