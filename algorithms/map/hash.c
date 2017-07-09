#include <stdio.h>
#include "./test.c"

unsigned int hash(char *string, int len) {
    unsigned int hash = 273;  

    for (int i = 0; i < len; ++i) {
      hash = hash * 33 + string[i];
    }

    return hash;
}

int main(int argc, char **argv) {
    char *str = "hello";
    unsigned int h = hash(str, 5);
    printf("%u", h);
    checkint(1, 2);
    return 0;
}
