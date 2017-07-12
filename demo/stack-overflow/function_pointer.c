#include <stdio.h>

void print() {
    printf("hello~\n");
}

void exec(void func()) {
    printf("exec: ");
    func();
}

void exec_ptr(void (*func)()) {
    printf("exec_ptr: ");
    (*func)();
}

int main() {
    // exec(print);
    exec_ptr(print);
}
