// Insertion Sort
// mtso 2017

#include <stdio.h>

void check(int got, int want) {
    if (got != want) {
        printf("FAILED %s:%d %d!=%d\n", __FILE__, __LINE__, got, want);
    } else {
        printf("passed\n");
    }
}

int desc(int a, int b) {
    return b - a;
}

int asc(int a, int b) {
    return a - b;
}

void insertion_sort(int *nums, int numlen, int (*comparer)(int, int)) {
    int temp;
    int curr;

    for (int i=0; i<numlen; ++i) {
        temp = *(nums+i);
        curr = i - 1;
        // while (curr >= 0 && temp > *(nums+curr)) {
        while (curr >= 0 && (*comparer)(temp, *(nums+curr)) < 0) {
            *(nums+curr+1) = *(nums+curr);
            curr = curr-1;
        }
        *(nums+curr+1) = temp;
    }
}

void printnums(int *nums, int numlen) {
    for (int i = 0; i < numlen; ++i) {
        printf("%d ", *(nums++));
    }
    printf("\n");
}

const int SIZE = 4;

int main(int argc, char** argv) {
    int nums[SIZE] = {3, 2, 4, 1};

    printnums(nums, SIZE);
    insertion_sort(nums, SIZE, asc);
    printnums(nums, SIZE);

    for (int i = 0; i < SIZE; ++i) {
        check(*(nums+i), 1 + i);
    }
}
