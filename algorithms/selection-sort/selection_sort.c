// Selection Sort
// mtso 2017

#include <stdio.h>

void check(int got, int want) {
    if (got != want) {
        printf("FAILED %s:%d %d!=%d\n", __FILE__, __LINE__, got, want);
    } else {
        printf("passed\n");
    }
}

void swap(int *a, int *b) {
    int temp = *a;
    *a = *b;
    *b = temp;
}

void selection_sort(int *nums, int len, int comparer(int, int)) {
    int selection;

    for (int i = 0; i < len; ++i) {
        selection = i;
        for (int j = i; j < len; ++j) {
            if (comparer(*(nums+j), *(nums+selection)) > 0) {
                selection = j;
            }
        }
        swap(nums+i, nums+selection);
    }
}

int desc(int a, int b) { return a - b; }
int asc(int a, int b) { return b - a; }

const int SIZE = 5;

void printnums(int *nums, int len) {
    for (int i = 0; i < len; ++i) {
        printf("%d ", *(nums+i));
    }
    printf("\n");
}

int main(int argc, char **argv) {
    int nums[SIZE] = {4, 3, 1, 5, 2};

    selection_sort(nums, SIZE, asc);
    printnums(nums, SIZE);

    for (int i = 0; i < SIZE; i++) {
        check(*(nums+i), i+1);
    }
}
