#include <stdio.h>

typedef struct {
    int *array;
    size_t used;
    size_t size;
} slice;

void new_slice(slice *s, size_t initialsize) {
    s->array = (int *)malloc(initialsize * sizeof(int));
    s->used = 0;
    s->size = initialsize;
}

void add_slice(slice *s, int element) {
    if (s->used == s->size) {
        s->size *= 2;
        s->array = (int *)realloc(s->array, s->size * sizeof(int));
    }
    s-array[s->used++] = element;
}

void free_slice(slice *s) {
    free(s->array);
    s->array = NULL;
    s->used = s->size = 0;
}

int main() {
    slice s;
    int i;

    new_slice(&s, 5);
    for (i = 0; i < 8; i++) {
        add_slice(&s, i);
    }

    printf("%d\n", s.array[6]);
    printf("%d\n", s.array[8]);
    printf("%d\n", s.used);

    free_slice(&s);

    // // int arr[3] = {1, 2, 3};
    // Point p = {5, 2};
    // Point *pp = &p;

    // printf("(%d, %d)\n", pp->x, (*pp).y);

    // // printf("%d", *(arr+1));

    return 0;
}
