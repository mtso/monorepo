#include <stdio.h>

/* check is an assertion function for two ints. */
int check(int want, int got) {
    if (want != got) {
        printf("FAIL %s:%d: %d != %d\n", __FILE__, __LINE__, want, got);
        return 0;
    }
    return 1;
}

/* swap swaps the values that a and b point to. */
void swap(int *a, int *b) {
    int temp = *a;
    *a = *b;
    *b = temp;
}

/* bubble_sort iteratively compares the current value
 * and the next value to sort an array in O(n^2) */
void bubble_sort(
    int *nums,
    int numlen,
    int (*comparer)(int, int)
) {
    for (int i = 0; i < numlen-1; ++i) {
        for (int j = i; j < numlen-1; ++j) {
            if ((*comparer)(*(nums+j), *(nums+j+1)) > 0){
                swap((nums+j), (nums+j+1));
            }
        }
    }
}

/* printarray outputs an array of ints to stdio. */
void printarray(int *nums, int numlen) {
    for (int i = 0; i < numlen; ++i) {
        printf("%d ", *(nums+i));
    }
}

/* asc compares a to b for ascending order. */
int asc(int a, int b) {
    return a - b;
}

/* desc compares a to b for descending order. */
int desc(int a, int b) {
    return b - a;
}

const int SIZE = 5;

int main() {
    int arr[SIZE] = {1, 7, 3, 2, 1};

    printarray(arr, SIZE);
    printf("\n");

    bubble_sort(arr, SIZE, desc);

    printarray(arr, SIZE);
    printf("\n");

    return 0;
}
