#include <stdio.h>

int check(int want, int got) {
    if (want != got) {
        printf("FAIL %s:%d: %d != %d\n", __FILE__, __LINE__, want, got);
        return 0;
    }
    return 1;
}

const int SIZE = 10;

int main() {
    int arr[SIZE] = {1, 2, 3, 4};

    if (!check(1, 2)) {
        return 1;
    }

    return 0;
}
