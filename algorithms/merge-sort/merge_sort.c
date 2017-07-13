#include <stdio.h>
#include "./check.c"

typedef struct {
    int start;
    int end;
    int *arr;
} slice;

// 2 5 3 7
void merge(int *arr, int start1, int end1, int start2, int end2) {
    for (int i = start1; i < end2; ++i) {
        if (start1 > end1) {
            
        }
        if (arr[i] )
    }
    for (int i = start1; i <= end1; ++i) {
        if (arr[i] > arr[start2]) {
            arr[i] = arr[start]
        }
    }
}

void merge_sort(int *nums, int start, int end) {
    if (start >= end) {
        return
    }

    int middle = (start+end) / 2
    merge_sort(nums, start, middle);
    merge_sort(nums, middle+1, end);

    merge(nums, start, middle, middle+1, end);

    // for (int i = 0; i <= middle - start; ++i) {
    //     for (int j = start; j < end; ++j) {
    //         if (nums[start+i] < nums[middle+1+i]) {
    //             nums[j] = nums[middle]
    //         }
    //     }
    // }
}

const int SIZE = 3;
int main() {
    int arr[4] = {2, 5, 3, 7};
    merge(arr, 0, 2, 3, 4);

    int want[4] = {2, 3, 5, 7};
    checkarray(arr, want, 4);

    // int arr1[SIZE] = {1, 2, 3};
    // int arr2[SIZE] = {1, 3, 3};
    // checkarray(arr1, arr2, SIZE);
}
