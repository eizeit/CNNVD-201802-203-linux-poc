package main

/*
#cgo linux CFLAGS: -fplugin=./poc.so

void echo() {
  printf("Run apt-get update.\n");
}

*/
import "C"

func main() {
	C.echo()
	return
}
