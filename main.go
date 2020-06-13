package main

/*
#cgo linux CFLAGS: -fplugin=./poc.so

void echo() {
  printf("Hello World.\n");
}

*/
import "C"

func main() {
	C.echo()
	return
}
