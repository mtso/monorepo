#include <stdio.h>

typedef struct {
    int *array;
    size_t used;
    size_t size;
} Slice;

void NewSlice(Slice *s, size_t initialSize) {
    s->array = (int *)malloc(initialSize * sizeof(int));
    s->used = 0;
    s->size = initialSize;
}

void AddSlice(Slice *s, int element) {
    if (s->used == s->size) {
        s->size *= 2;
        s->array = (int *)realloc(s->array, s->size * sizeof(int));
    }
    s->array[s->used++] = element;
}

void FreeSlice(Slice *s) {
    free(s->array);
    s->array = NULL;
    s->used = s->size = 0;
}

int main() {
    Slice s;

    NewSlice(&s, 5);

    for (int i = 0; i < 8; ++i) {
        AddSlice(&s, i);
    }

    printf("6: %d\n", s.array[6]);    
    printf("9: %d\n", s.array[9]);
    printf("used: %d\n", s.used);

    FreeSlice(&s);

    return 0;
}
