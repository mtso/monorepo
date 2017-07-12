#include <stdio.h>
#include "./check.c"

const int SIZE = 3;
int main() {
    int arr1[SIZE] = {1, 2, 3};
    int arr2[SIZE] = {1, 3, 3};
    checkarray(arr1, arr2, SIZE);
}
