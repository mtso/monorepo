#include <stdio.h>

void print() {
    printf("hello~\n");
}

void exec(void func()) {
    func();
}

int main() {
    exec(print);
}
