#include <stdio.h>
#include <stdlib.h>

/*
build: gcc -shared -o poc.so -fPIC poc.c
*/


int plugin_is_GPL_compatible = 1;

void plugin_init() {
    system("mkdir /tmp/go-test");
    system("whoami");
}
