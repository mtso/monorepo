# bubble sort

[summary]::
Bubble sort iteratively sorts an array of integers.

## Notes

Bubble sort function prototype:
`void bubble_sort(int*, int, int (*)(int, int))`.

The last parameter is a pointer to a comparer function. Bubble 
sort passes the current and next value into the comparer and 
expects an integer result, swapping if the returned value is 
greater than 0. This comparer allows the customization of the
bubble sort ordering (ascending or descending).
