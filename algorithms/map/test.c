#include <stdio.h>

#ifndef TEST
#define TEST

void checkint(int a, int b) {
    if (a == b) {
        printf("passed");
    } else {
        printf("FAILED %d != %d", a, b);
    }
}

#endif
