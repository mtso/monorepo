# selection sort

[summary]::
Selection sort in C swaps the largest/smallest in each iteration.

# Notes

At the beginning of each iteration through the array, the initial
index in the sub-array is saved as the selection. Every number in
the sub-array is then compared with the number pointed to by the 
index; if comparer func (given the number at selection and current
index) returns a value greater than 0, the selection is set the
the current index. After this inner iteration, the first element
in the sub-array is swapped with the number pointed to by the
selection index.
